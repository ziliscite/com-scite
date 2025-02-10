package repository

import (
	"context"
	"database/sql"
	"errors"
	"fmt"

	"github.com/jackc/pgconn"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/ziliscite/micro-auth/auth/internal/domain"
)

var (
	ErrEditConflict   = errors.New("conflict")
	ErrRecordNotFound = errors.New("not found")
	ErrDuplicateEntry = errors.New("duplicate")
)

type UserRepository interface {
	GetByEmail(ctx context.Context, email string) (*domain.User, error)
	Insert(ctx context.Context, user *domain.User) error

	GetById(ctx context.Context, id int64) (*domain.User, error)
	Update(ctx context.Context, user *domain.User) error
}

type userRepo struct {
	db *pgxpool.Pool
}

func NewUserRepository(db *pgxpool.Pool) UserRepository {
	return &userRepo{db: db}
}

func (u userRepo) GetByEmail(ctx context.Context, email string) (*domain.User, error) {
	query := `
        SELECT id, username, email, password_hash, activated, created_at
        FROM users
        WHERE email = $1;
	`

	var user domain.User

	var hash []byte
	err := u.db.QueryRow(ctx, query, email).Scan(
		&user.ID, &user.Username, &user.Email,
		&hash, &user.Activated, &user.CreatedAt,
	)

	user.Password.InsertHash(hash)

	if err != nil {
		switch {
		case errors.Is(err, sql.ErrNoRows):
			return nil, ErrRecordNotFound
		default:
			return nil, fmt.Errorf("something's wrong: %w", err)
		}
	}

	return &user, nil
}

func (u userRepo) Insert(ctx context.Context, user *domain.User) error {
	query := `
        INSERT INTO users (username, email, password_hash, activated) 
        VALUES ($1, $2, $3, $4)
        RETURNING id, created_at, version;
	`

	args := []any{user.Username, user.Email, user.Hash(), user.Activated}

	err := u.db.QueryRow(ctx, query, args...).Scan(&user.ID, &user.CreatedAt, &user.Version)
	if err != nil {
		var pgErr *pgconn.PgError
		switch {
		case errors.As(err, &pgErr) && pgErr.Code == "23505":
			return ErrDuplicateEntry
		case errors.Is(err, sql.ErrNoRows):
			return ErrEditConflict
		default:
			return fmt.Errorf("something's wrong: %w", err)
		}
	}

	return nil
}

func (u userRepo) GetById(ctx context.Context, id int64) (*domain.User, error) {
	query := `
        SELECT id, username, email, password_hash, activated, created_at, updated_at, version
        FROM users
        WHERE id = $1;
	`

	var user domain.User

	var hash []byte
	if err := u.db.QueryRow(ctx, query, id).Scan(
		&user.ID, &user.Username, &user.Email,
		&hash, &user.Activated, &user.CreatedAt,
		&user.UpdatedAt, &user.Version,
	); err != nil {
		switch {
		case errors.Is(err, sql.ErrNoRows):
			return nil, ErrRecordNotFound
		default:
			return nil, fmt.Errorf("something's wrong: %w", err)
		}
	}

	user.Password.InsertHash(hash)

	return &user, nil
}

func (u userRepo) Update(ctx context.Context, user *domain.User) error {
	query := `
        UPDATE users 
        SET username = $1, email = $2, password_hash = $3, activated = $4, version = version + 1, updated_at = NOW()
        WHERE id = $5 AND version = $6
        RETURNING version, updated_at;
	`

	args := []any{user.Username, user.Email, user.Hash(), user.Activated, user.ID, user.Version}

	err := u.db.QueryRow(ctx, query, args...).Scan(&user.Version, &user.UpdatedAt)
	if err != nil {
		switch {
		case errors.Is(err, sql.ErrNoRows):
			return ErrEditConflict
		default:
			return fmt.Errorf("something's wrong: %w", err)
		}
	}

	return nil
}
