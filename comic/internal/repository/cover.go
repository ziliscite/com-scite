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
	New(ctx context.Context, comicId int64, cover *domain.Cover) error
	Deactivate(ctx context.Context, comicId int64) error
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
func (r *coverRepository) New(ctx context.Context, comicId int64, cover *domain.Cover) error {
	query := `
		INSERT INTO cover(comic_id, url) 
		VALUES ($1, $2)
	`

	_, err := r.db.Exec(ctx, query, comicId, cover.URL)
	if err != nil {
		var pgErr *pgconn.PgError
		switch {
		case errors.As(err, &pgErr) && pgErr.Code == "23505":
			return fmt.Errorf("%w: %s already exist", ErrDuplicate, cover.URL)
		default:
			return err
		}
	}

	return nil
}

func (r *coverRepository) Deactivate(ctx context.Context, comicId int64) error {
	query := `
		UPDATE cover
		SET is_current = false, updated_at = now()
		WHERE comic_id = $1 AND is_current = true;
    `

	_, err := r.db.Exec(ctx, query, comicId)
	if err != nil {
		switch {
		case errors.Is(err, pgx.ErrNoRows):
			return ErrNotFound
		default:
			return err
		}
	}

	return nil
}

func (r *coverRepository) GetAll(ctx context.Context, comicId int64) ([]domain.Cover, error) {
	query := `
		SELECT cover_id, comic_id, url, is_current, created_at, updated_at 
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
		SELECT cover_id, comic_id, url, is_current, created_at, updated_at 
		FROM cover 
		WHERE comic_id = $1 AND is_current = true
		ORDER BY created_at;
 	`

	var cover domain.Cover
	if err := r.db.QueryRow(ctx, query, comicId).Scan(
		&cover.ID, &cover.ComicID,
		&cover.URL, &cover.IsCurrent,
		&cover.CreatedAt,
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
