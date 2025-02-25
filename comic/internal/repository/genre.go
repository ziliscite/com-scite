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

type GenreRepository interface {
	// MassUpsert will do batch get or insert genres
	MassUpsert(ctx context.Context, genres []*domain.Genre) error
	New(ctx context.Context, genre *domain.Genre) error

	GetAll(ctx context.Context) ([]domain.Genre, error)

	Edit(ctx context.Context, genre *domain.Genre) error
	Delete(ctx context.Context, genreId int64) error
}

type genreRepository struct {
	db *pgxpool.Pool
}

func NewGenreRepository(db *pgxpool.Pool) GenreRepository {
	return &genreRepository{db: db}
}

func (g *genreRepository) MassUpsert(ctx context.Context, genres []*domain.Genre) error {
	batch := &pgx.Batch{}
	for _, gr := range genres {
		batch.Queue(`
			INSERT INTO genre (name)
			VALUES ($1)
			ON CONFLICT (name) DO UPDATE SET name=excluded.name
			RETURNING name, genre_id
		`, gr.Name)
	}

	br := g.db.SendBatch(ctx, batch)
	defer br.Close()

	for _, gr := range genres {
		if err := br.QueryRow().Scan(&gr.Name, &gr.ID); err != nil {
			return err
		}
	}

	return nil
}

func (g *genreRepository) New(ctx context.Context, genre *domain.Genre) error {
	query := `
		INSERT INTO genre (name) VALUES ($1) RETURNING genre_id;
	`

	if err := g.db.QueryRow(ctx, query, genre.ID).Scan(&genre.ID); err != nil {
		var pgErr *pgconn.PgError
		switch {
		case errors.As(err, &pgErr) && pgErr.Code == "23505":
			return fmt.Errorf("%w: %s already exist", ErrDuplicate, genre.Name)
		default:
			return err
		}
	}

	return nil
}

func (g *genreRepository) GetAll(ctx context.Context) ([]domain.Genre, error) {
	query := `
		SELECT name, genre_id FROM genre;
	`

	rows, err := g.db.Query(ctx, query)
	if err != nil {
		switch {
		case errors.Is(err, pgx.ErrNoRows):
			return nil, ErrNotFound
		default:
			return nil, err
		}
	}
	defer rows.Close()

	genres := make([]domain.Genre, 0)
	for rows.Next() {
		var genre domain.Genre
		if err := rows.Scan(&genre.Name, &genre.ID); err != nil {
			return nil, err
		}
		genres = append(genres, genre)
	}

	return genres, nil
}

func (g *genreRepository) Edit(ctx context.Context, genre *domain.Genre) error {
	query := `
		UPDATE genre SET name = $1 WHERE genre_id = $2;
	`

	_, err := g.db.Exec(ctx, query, genre.Name, genre.ID)
	if err != nil {
		var pgErr *pgconn.PgError
		switch {
		case errors.Is(err, pgx.ErrNoRows):
			return ErrNotFound
		case errors.As(err, &pgErr) && pgErr.Code == "23505":
			return fmt.Errorf("%w: %s already exist", ErrDuplicate, genre.Name)
		default:
			return nil
		}
	}

	return nil
}

func (g *genreRepository) Delete(ctx context.Context, genreId int64) error {
	query := `
		DELETE FROM genre WHERE genre_id = $1;
	`

	cmd, err := g.db.Exec(ctx, query, genreId)
	if err != nil {
		return err
	}

	if cmd.RowsAffected() == 0 {
		return ErrNotFound
	}

	return nil
}
