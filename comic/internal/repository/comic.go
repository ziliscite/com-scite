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

type ComicRepository interface {
	Create(ctx context.Context, comic *domain.Comic) error

	Index(ctx context.Context) ([]domain.Comic, error)
	Get(ctx context.Context, comicId int64) (*domain.Comic, error)
	GetBySlug(ctx context.Context, comicSlug string) (*domain.Comic, error)

	Update(ctx context.Context, comic *domain.Comic) error
	Delete(ctx context.Context, comicId int64) error
}

type comicRepository struct {
	db *pgxpool.Pool
}

func NewComicRepository(db *pgxpool.Pool) ComicRepository {
	return &comicRepository{
		db: db,
	}
}

func (c *comicRepository) Create(ctx context.Context, comic *domain.Comic) error {
	query := `
		INSERT INTO comic(title, slug, description, author, artist, status, type) 
		VALUES ($1, $2, $3, $4, $5, $6, $7)
		RETURNING comic_id;
	`

	args := []any{comic.Title, comic.Slug, comic.Description, comic.Author, comic.Artist, comic.Status.String(), comic.Type.String()}
	if err := c.db.QueryRow(ctx, query, args...).Scan(&comic.ID); err != nil {
		var pgErr *pgconn.PgError
		switch {
		case errors.As(err, &pgErr) && pgErr.Code == "23505":
			return fmt.Errorf("%w: %s already exist", ErrDuplicate, comic.Title)
		default:
			return err
		}
	}

	return nil
}

func (c *comicRepository) Index(ctx context.Context) ([]domain.Comic, error) {
	//TODO implement me
	panic("implement me")
}

func (c *comicRepository) Get(ctx context.Context, comicId int64) (*domain.Comic, error) {
	query := `
		SELECT
			c.comic_id, c.title, c.slug, c.description,
			c.author, c.artist, c.status, c.type,
			ARRAY_AGG(g.name ORDER BY g.name) AS genre,
			c.created_at, c.version
		FROM comic c
		JOIN comic_genre cg ON c.comic_id = cg.comic_id
		JOIN genre g ON cg.genre_id = g.genre_id
		WHERE c.comic_id = $1
		GROUP BY c.comic_id;
	`

	var comic domain.Comic
	if err := c.db.QueryRow(ctx, query, comicId).Scan(
		&comic.ID, &comic.Title, &comic.Slug,
		&comic.Description, &comic.Author,
		&comic.Artist, &comic.Status, &comic.Type,
		&comic.Genres, &comic.CoverUrl,
		&comic.CreatedAt, &comic.UpdatedAt, &comic.Version,
	); err != nil {
		switch {
		case errors.Is(err, pgx.ErrNoRows):
			return nil, ErrNotFound
		default:
			return nil, err
		}
	}

	return &comic, nil
}

func (c *comicRepository) GetBySlug(ctx context.Context, comicSlug string) (*domain.Comic, error) {
	query := `
		SELECT
			c.comic_id, c.title, c.slug, c.description,
			c.author, c.artist, c.status, c.type,
			ARRAY_AGG(g.name ORDER BY g.name) AS genre,
			c.created_at, c.version
		FROM comic c
		JOIN comic_genre cg ON c.comic_id = cg.comic_id
		JOIN genre g ON cg.genre_id = g.genre_id
		WHERE c.slug = $1
		GROUP BY c.comic_id;
	`

	var comic domain.Comic
	if err := c.db.QueryRow(ctx, query, comicSlug).Scan(
		&comic.ID, &comic.Title, &comic.Slug,
		&comic.Description, &comic.Author,
		&comic.Artist, &comic.Status, &comic.Type,
		&comic.Genres, &comic.CoverUrl,
		&comic.CreatedAt, &comic.UpdatedAt, &comic.Version,
	); err != nil {
		switch {
		case errors.Is(err, pgx.ErrNoRows):
			return nil, fmt.Errorf("comic is %w", ErrNotFound)
		default:
			return nil, err
		}
	}

	return &comic, nil
}

// Update can only update comic instance, not genre and cover
func (c *comicRepository) Update(ctx context.Context, comic *domain.Comic) error {
	query := `
		UPDATE comic 
		SET title = $1, type = $2, slug = $3, description = $4, 
		    author = $5, artist = $6, status = $7, type = $8, 
		    updated_at = now(), version = version + 1
		WHERE comic_id = $9 AND version = $10
		RETURNING version
	`

	args := []any{
		comic.Title, comic.Slug, comic.Description,
		comic.Author, comic.Artist, comic.Status.String(),
		comic.Type.String(), comic.ID, comic.Version,
	}

	if err := c.db.QueryRow(ctx, query, args...).Scan(&comic.Version); err != nil {
		var pgErr *pgconn.PgError
		switch {
		case errors.As(err, &pgErr) && pgErr.Code == "23505":
			return fmt.Errorf("%w: %s already exist", ErrDuplicate, comic.Title)
		// Couldn't be a `not found` error, since we GET the comic FIRST in the service layer
		// When it is `not found`, it means, a version is different, hence error conflict
		case errors.Is(err, pgx.ErrNoRows):
			return ErrConflict
		default:
			return err
		}
	}

	return nil
}

// Delete will only delete the instance of a comic and its cover.
// Doesn't have to handle the delete on its own, since on delete cascade will happen
func (c *comicRepository) Delete(ctx context.Context, comicId int64) error {
	query := `
		DELETE FROM comic WHERE comic_id = $1
	`

	cmd, err := c.db.Exec(ctx, query, comicId)
	if err != nil {
		return err
	}

	if cmd.RowsAffected() == 0 {
		return ErrNotFound
	}

	return nil
}
