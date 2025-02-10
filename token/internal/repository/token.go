package repository

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/ziliscite/micro-auth/token/internal/domain"
)

var ErrRecordNotFound = errors.New("record not found")

type TokenRepository interface {
	Insert(ctx context.Context, token *domain.Token) error
	DeleteAllForUser(ctx context.Context, userID int64) error
	GetUserId(ctx context.Context, tokenHash []byte) (int64, error)
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
        INSERT INTO tokens (user_id, token_hash, expired_at) VALUES ($1, $2, $3);
	`

	args := []any{token.UserID, token.Hash, token.Expiry}

	_, err := t.db.Exec(ctx, query, args...)
	return err
}

// GetUserId get user id from token
func (t tokenRepository) GetUserId(ctx context.Context, tokenHash []byte) (int64, error) {
	query := `
        SELECT user_id FROM tokens WHERE token_hash = $1;
	`

	var userId int64
	if err := t.db.QueryRow(ctx, query, tokenHash).Scan(&userId); err != nil {
		switch {
		case errors.Is(err, sql.ErrNoRows):
			return 0, ErrRecordNotFound
		default:
			return 0, fmt.Errorf("something's wrong: %w", err)
		}
	}

	return userId, nil
}

// DeleteAllForUser deletes all tokens for a specific user when their account has been activated.
func (t tokenRepository) DeleteAllForUser(ctx context.Context, userID int64) error {
	query := `
        DELETE FROM tokens WHERE user_id = $1;
	`

	_, err := t.db.Exec(ctx, query, userID)
	return err
}
