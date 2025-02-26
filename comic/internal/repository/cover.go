package repository

import (
	"context"
	"errors"
	"fmt"
	"github.com/jackc/pgconn"
	"github.com/jackc/pgx/v4"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/ziliscite/com-scite/comic/internal/domain"
)

type CoverRepository interface {
	New(ctx context.Context, cover *domain.Cover) (string, error)
	ReActive(ctx context.Context, comicId int64, fileKey string) error

	GetActive(ctx context.Context, comicId int64) (*domain.Cover, error)
	GetAll(ctx context.Context, comicId int64) ([]domain.Cover, error)
}

type coverRepository struct {
	db *pgxpool.Pool
}

func NewCoverRepository(db *pgxpool.Pool) CoverRepository {
	return &coverRepository{db: db}
}

// New will insert a new cover, set default to true, and disable old covers
// Or does that in the service?
func (r *coverRepository) New(ctx context.Context, cover *domain.Cover) (string, error) {
	tx, err := r.db.Begin(ctx)
	if err != nil {
		return "", err
	}
	defer tx.Rollback(ctx)

	var oldKey string
	if err = tx.QueryRow(ctx, `
		UPDATE cover
		SET is_current = false, updated_at = now()
		WHERE comic_id = $1 AND is_current = true
		RETURNING file_key
    `, cover.ComicID).Scan(&oldKey); err != nil {
		switch {
		case errors.Is(err, pgx.ErrNoRows):
			// ignore --
		default:
			return "", err
		}
	}

	if _, err = tx.Exec(ctx, `
		INSERT INTO cover(comic_id, file_key) 
		VALUES ($1, $2)
	`, cover.ComicID, cover.FileKey); err != nil {
		var pgErr *pgconn.PgError
		switch {
		case errors.As(err, &pgErr) && pgErr.Code == "23505":
			return "", fmt.Errorf("%w: %s already exist", ErrDuplicate, cover.FileKey)
		default:
			return "", err
		}
	}

	return oldKey, tx.Commit(ctx)
}

func (r *coverRepository) ReActive(ctx context.Context, comicId int64, fileKey string) error {
	tx, err := r.db.Begin(ctx)
	if err != nil {
		return err
	}
	defer tx.Rollback(ctx)

	if _, err = tx.Exec(ctx, `
		UPDATE cover
		SET is_current = false, updated_at = now()
		WHERE comic_id = $1 AND is_current = true
    `, comicId); err != nil {
		switch {
		case errors.Is(err, pgx.ErrNoRows):
			return ErrNotFound
		default:
			return err
		}
	}

	if _, err = tx.Exec(ctx, `
		UPDATE cover
		SET is_current = true, updated_at = now()
		WHERE comic_id = $1 AND file_key = $2
	`, comicId, fileKey); err != nil {
		switch {
		case errors.Is(err, pgx.ErrNoRows):
			return ErrNotFound
		default:
			return err
		}
	}

	return tx.Commit(ctx)
}

func (r *coverRepository) GetAll(ctx context.Context, comicId int64) ([]domain.Cover, error) {
	query := `
		SELECT cover_id, comic_id, file_key, is_current, created_at, updated_at 
		FROM cover 
		WHERE comic_id = $1;
 	`

	rows, err := r.db.Query(ctx, query, comicId)
	if err != nil {
		switch {
		case errors.Is(err, pgx.ErrNoRows):
			return nil, ErrNotFound
		default:
			return nil, err
		}
	}
	defer rows.Close()

	covers := make([]domain.Cover, 0)
	for rows.Next() {
		var cover domain.Cover
		if err = rows.Scan(&cover); err != nil {
			return nil, err
		}
		covers = append(covers, cover)
	}

	return covers, nil
}

func (r *coverRepository) GetActive(ctx context.Context, comicId int64) (*domain.Cover, error) {
	// Get the latest cover
	query := `
		SELECT cover_id, comic_id, file_key, is_current, created_at, updated_at 
		FROM cover 
		WHERE comic_id = $1 AND is_current = true;
 	`

	var cover domain.Cover
	if err := r.db.QueryRow(ctx, query, comicId).Scan(
		&cover.ID, &cover.ComicID, &cover.FileKey,
		&cover.IsCurrent, &cover.CreatedAt, &cover.UpdatedAt,
	); err != nil {
		switch {
		case errors.Is(err, pgx.ErrNoRows):
			return nil, ErrNotFound
		default:
			return nil, err
		}
	}

	return &cover, nil
}
