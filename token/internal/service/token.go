package service

import (
	"context"
	"errors"
	"github.com/ziliscite/micro-auth/token/internal/domain"
	"github.com/ziliscite/micro-auth/token/internal/repository"
	"time"
)

var (
	ErrInvalidToken = errors.New("invalid token")
)

type TokenService interface {
	New(ctx context.Context, userID int64, ttl time.Duration) (*domain.Token, error)
}

type tokenService struct {
	tr repository.TokenRepository
}

func NewTokenService(tr repository.TokenRepository) TokenService {
	return &tokenService{
		tr: tr,
	}
}

func (t tokenService) New(ctx context.Context, userID int64, ttl time.Duration) (*domain.Token, error) {
	token, err := domain.GenerateToken(userID, ttl)
	if err != nil {
		return nil, err
	}

	err = t.tr.Insert(ctx, token)
	if err != nil {
		return nil, err
	}

	return token, nil
}

func (t tokenService) Activate(ctx context.Context, tokenString string) (int64, error) {
	err := domain.ValidateTokenPlaintext(tokenString)
	if err != nil {
		return 0, ErrInvalidToken
	}

	tokenHash := domain.GetTokenHash(tokenString)

	return t.tr.GetUserId(ctx, tokenHash)
}
