package repository

import (
	"context"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/ziliscite/micro-auth/token/internal/domain"
)

type TokenRepository interface {
	Insert(ctx context.Context, token *domain.Token) error
	DeleteAllForUser(ctx context.Context, userID int64) error
}

type tokenRepository struct {
	db *pgxpool.Pool
}

func NewTokenRepository(db *pgxpool.Pool) TokenRepository {
	return tokenRepository{
		db: db,
	}
}

// Insert adds the data for a specific token to the tokens table.
func (t tokenRepository) Insert(ctx context.Context, token *domain.Token) error {
	query := `
        INSERT INTO tokens (hash, user_id, expiry) 
        VALUES ($1, $2, $3)
	`

	args := []any{token.Hash, token.UserID, token.Expiry}

	_, err := t.db.Exec(ctx, query, args...)
	if err != nil {
		return err
	}

	return nil
}

// DeleteAllForUser deletes all tokens for a specific user when their account has been activated.
func (t tokenRepository) DeleteAllForUser(ctx context.Context, userID int64) error {
	query := `
        DELETE FROM tokens 
        WHERE user_id = $1
	`

	_, err := t.db.Exec(ctx, query, userID)
	if err != nil {
		return err
	}

	return nil
}
