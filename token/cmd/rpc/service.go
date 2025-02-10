package main

import (
	"context"
	"errors"
	"github.com/ziliscite/micro-auth/token/internal/domain"
	"github.com/ziliscite/micro-auth/token/internal/service"
	pb "github.com/ziliscite/micro-auth/token/pkg/protobuf"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"log/slog"
	"time"
)

type Service struct {
	pb.UnimplementedActivationServiceServer
	ts  service.TokenService
	as  pb.AuthServiceClient
	pub service.MailPublisher
}

func newClient(ts service.TokenService, as pb.AuthServiceClient, pub service.MailPublisher) *Service {
	return &Service{ts: ts, as: as, pub: pub}
}

func (s Service) CreateActivation(ctx context.Context, req *pb.ActivationRequest) (*pb.ActivationResponse, error) {
	ctx, cancel := context.WithTimeout(ctx, 3*time.Second)
	defer cancel()

	token, err := s.ts.New(ctx, req.GetUserId(), time.Duration(24)*time.Hour)
	if err != nil {
		slog.Error("CreateActivation failed", "error", err.Error())
		switch {
		case errors.Is(err, context.DeadlineExceeded):
			return nil, status.Error(codes.DeadlineExceeded, err.Error())
		case errors.Is(err, context.Canceled):
			return nil, status.Error(codes.Canceled, err.Error())
		}
	}

	if err = s.pub.SendMail(domain.Mail{
		ID:        token.UserID,
		Username:  req.GetUsername(),
		Email:     req.GetEmail(),
		Token:     token.Plaintext,
		ExpiresAt: token.Expiry,
	}); err != nil {
		slog.Error("CreateActivation failed", "error", err.Error())
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &pb.ActivationResponse{
		Status: "activation email has been sent!",
	}, nil
}

func (s Service) ActivateUser(ctx context.Context, req *pb.ActivateRequest) (*pb.ActivateResponse, error) {
	ctx, cancel := context.WithTimeout(ctx, 3*time.Second)
	defer cancel()

	userId, err := s.ts.Activate(ctx, req.GetTokenString())
	if err != nil {
		slog.Error("Activation failed", "error", err.Error())
		switch {
		case errors.Is(err, context.DeadlineExceeded):
			return nil, status.Error(codes.DeadlineExceeded, err.Error())
		case errors.Is(err, context.Canceled):
			return nil, status.Error(codes.Canceled, err.Error())
		case errors.Is(err, service.ErrInvalidToken):
			return nil, status.Error(codes.InvalidArgument, err.Error())
		}
	}

	// call an auth client to update the user
	//
	resp, err := s.as.ActivateUser(ctx, &pb.ActivateUserRequest{
		UserId: userId,
	})
	if err != nil {
		slog.Error("Activate failed", "error", err.Error())
		return nil, status.Error(codes.Internal, err.Error())
	}
	if !resp.Activated {
		return nil, status.Error(codes.FailedPrecondition, "user has not been activated")
	}

	user := resp.GetResponse()

	// send congrats mail
	if err = s.pub.SendMail(domain.Mail{
		ID:        -1,
		Username:  user.GetUsername(),
		Email:     user.GetEmail(),
		Token:     "To simulate congrats email",
		ExpiresAt: time.Now(),
	}); err != nil {
		slog.Error("Send welcome email failed", "error", err.Error())
		// not return, aint that important anyway
	}

	return &pb.ActivateResponse{
		Activated: resp.Activated,
		Status:    "User has been activated!",
	}, nil
}
