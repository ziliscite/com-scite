package main

import (
	"context"
	"errors"
	"github.com/ziliscite/micro-auth/auth/internal/repository"
	"github.com/ziliscite/micro-auth/auth/internal/service"
	pb "github.com/ziliscite/micro-auth/auth/pkg/protobuf"
	"github.com/ziliscite/micro-auth/auth/pkg/token"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"time"
)

type Service struct {
	pb.UnimplementedAuthServiceServer
	us      service.UserService
	secrets string
}

func NewService(us service.UserService, secrets string) *Service {
	return &Service{
		us:      us,
		secrets: secrets,
	}
}

func (s Service) Register(ctx context.Context, req *pb.RegisterRequest) (*pb.RegisterResponse, error) {
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	user, err := s.us.SignUp(ctx, req.GetUsername(), req.GetEmail(), req.GetPassword())
	if err != nil {
		switch {
		case errors.Is(err, context.DeadlineExceeded):
			return nil, status.Error(codes.DeadlineExceeded, err.Error())
		case errors.Is(err, context.Canceled):
			return nil, status.Error(codes.Canceled, err.Error())
		case errors.Is(err, service.ErrInvalidUser):
			return nil, status.Error(codes.InvalidArgument, err.Error())
		case errors.Is(err, repository.ErrDuplicateEntry):
			return nil, status.Error(codes.AlreadyExists, err.Error())
		case errors.Is(err, repository.ErrEditConflict):
			return nil, status.Error(codes.FailedPrecondition, err.Error())
		default:
			return nil, status.Error(codes.Internal, err.Error())
		}
	}

	// call token service to create a new activation token
	//
	// in token service, asynchronously calls for the mailer service
	// to send the email notification with token and user information
	//

	return &pb.RegisterResponse{
		Response: &pb.User{
			Id:       user.ID,
			Username: user.Username,
			Email:    user.Email,
		},
	}, nil
}

func (s Service) Login(ctx context.Context, req *pb.LoginRequest) (*pb.LoginResponse, error) {
	ctx, cancel := context.WithTimeout(ctx, 6*time.Second)
	defer cancel()

	user, err := s.us.SignIn(ctx, req.GetEmail(), req.GetPassword())
	if err != nil {
		switch {
		case errors.Is(err, context.DeadlineExceeded):
			return nil, status.Error(codes.DeadlineExceeded, err.Error())
		case errors.Is(err, context.Canceled):
			return nil, status.Error(codes.Canceled, err.Error())
		case errors.Is(err, service.ErrInvalidCredentials):
			return nil, status.Error(codes.Unauthenticated, err.Error())
		default:
			return nil, status.Error(codes.Internal, err.Error())
		}
	}

	accessToken, exp, err := token.Create(user.ID, user.Activated, user.Email, s.secrets)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &pb.LoginResponse{
		Response: &pb.Token{
			Token:      accessToken,
			Expiration: exp.Unix(),
		},
	}, nil
}
