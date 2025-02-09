package service

import (
	"context"
	"github.com/ziliscite/micro-auth/token/internal/domain"
	"github.com/ziliscite/micro-auth/token/internal/repository"
	"time"
)

type TokenService interface {
	New(ctx context.Context, userID int64, ttl time.Duration) (*domain.Token, error)
}

type tokenService struct {
	tr repository.TokenRepository
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
