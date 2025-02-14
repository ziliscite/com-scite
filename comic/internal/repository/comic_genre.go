package repository

import (
	"context"
	"errors"
	"fmt"
	"github.com/jackc/pgconn"
	"github.com/jackc/pgx/v4"
	"github.com/jackc/pgx/v4/pgxpool"
)

type ComicGenreRepository interface {
	MassUpsert(ctx context.Context, comicId int64, genresId []int64) error
	Add(ctx context.Context, comicId int64, genreId int64) error
	Remove(ctx context.Context, comicId int64, genreId int64) error
}

type comicGenreRepository struct {
	db *pgxpool.Pool
}

func NewComicGenreRepository(db *pgxpool.Pool) ComicGenreRepository {
	return &comicGenreRepository{db: db}
}

func (c *comicGenreRepository) MassUpsert(ctx context.Context, comicId int64, genresId []int64) error {
	query := `
		INSERT INTO comicgenre (comic_id, genre_id) VALUES ($1, $2);
	`

	batch := &pgx.Batch{}
	for _, gr := range genresId {
		batch.Queue(query, comicId, gr)
	}

	cgr := c.db.SendBatch(ctx, batch)
	defer cgr.Close()

	for i := 0; i < len(genresId); i++ {
		_, err := cgr.Exec()
		if err != nil {
			return fmt.Errorf("failed to insert genre id %d for comic id %d: %v", genresId[i], comicId, err)
		}
	}

	return nil
}

func (c *comicGenreRepository) Add(ctx context.Context, comicId int64, genreId int64) error {
	query := `
		INSERT INTO comicgenre (comic_id, genre_id) VALUES ($1, $2);
	`

	_, err := c.db.Exec(ctx, query, comicId, genreId)
	if err != nil {
		var pgErr *pgconn.PgError
		switch {
		case errors.As(err, &pgErr) && pgErr.Code == "23505":
			return fmt.Errorf("%w: %v and %v relation already exist", ErrDuplicate, comicId, genreId)
		default:
			return err
		}
	}

	return nil
}

func (c *comicGenreRepository) Remove(ctx context.Context, comicId int64, genreId int64) error {
	query := `
		DELETE FROM comicgenre WHERE comic_id = $1 AND genre_id = $2;
	`

	cmd, err := c.db.Exec(ctx, query, comicId, genreId)
	if err != nil {
		return err
	}

	if cmd.RowsAffected() == 0 {
		return ErrNotFound
	}

	return nil
}
