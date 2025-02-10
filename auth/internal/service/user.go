package service

import (
	"context"
	"errors"
	"fmt"
	"github.com/ziliscite/micro-auth/auth/internal/domain"
	"github.com/ziliscite/micro-auth/auth/internal/repository"
	"github.com/ziliscite/micro-auth/auth/pkg/validator"
	"log/slog"
)

var (
	ErrInternal           = errors.New("internal error")
	ErrInvalidUser        = errors.New("invalid user")
	ErrInvalidCredentials = errors.New("invalid credentials")
)

type UserService interface {
	SignIn(ctx context.Context, email, password string) (*domain.User, error)
	SignUp(ctx context.Context, username, email, password string) (*domain.User, error)
}

type userServ struct {
	ur repository.UserRepository
}

func NewUserService(ur repository.UserRepository) UserService {
	return &userServ{ur}
}

func (u userServ) SignIn(ctx context.Context, email, password string) (*domain.User, error) {
	user, err := u.ur.GetByEmail(ctx, email)
	if err != nil {
		switch {
		case errors.Is(err, repository.ErrRecordNotFound):
			return nil, ErrInvalidCredentials
		default:
			return nil, fmt.Errorf("something's wrong: %w", err)
		}
	}

	ok, err := user.Password.Matches(password)
	if err != nil {
		return nil, err
	}
	if !ok {
		return nil, ErrInvalidCredentials
	}

	return user, nil
}

func (u userServ) SignUp(ctx context.Context, username, email, password string) (*domain.User, error) {
	user := domain.RegisterUser(username, email)

	err := user.Password.Set(password)
	if err != nil {
		return nil, ErrInternal
	}

	v := validator.New()

	user.Validate(v)
	if !v.Valid() {
		slog.Error("validation failed", "errors", v.Errors(), slog.String("username", username), slog.String("email", email))
		return nil, ErrInvalidUser
	}

	err = u.ur.Insert(ctx, user)
	if err != nil {
		return nil, err
	}

	return user, nil
}

// After activate, send congrats email

func (u userServ) Activate(ctx context.Context, id int64) (*domain.User, error) {
	user, err := u.ur.GetById(ctx, id)
	if err != nil {
		return nil, err
	}

	user.Activated = true

	err = u.ur.Update(ctx, user)
	if err != nil {
		return nil, err
	}

	return user, nil
}
