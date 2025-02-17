package service

import (
	"context"
	"fmt"
	"github.com/ziliscite/com-scite/comic/internal/domain"
	"github.com/ziliscite/com-scite/comic/internal/repository"
)

// Service layer is turning request from where ever it is, into the response
// Not tied to domain, at all, it seems.

type ComicService interface {
	Index(ctx context.Context) ([]domain.Comic, error)
	GetComicByID(ctx context.Context, comicId int64) (*domain.Comic, error)
	NewComic(ctx context.Context, title, description, author, artist, comicStatus, comicType string, genres []string) (*domain.Comic, error)
}

type comicService struct {
	cr  repository.ComicRepository
	cvr repository.CoverRepository
	gr  repository.GenreRepository
	cgr repository.ComicGenreRepository
}

func NewComicService(cr repository.ComicRepository, cvr repository.CoverRepository, gr repository.GenreRepository, cgr repository.ComicGenreRepository) ComicService {
	return &comicService{cr: cr, cvr: cvr, gr: gr, cgr: cgr}
}

func (c *comicService) Index(ctx context.Context) ([]domain.Comic, error) {
	return c.cr.Index(ctx)
}

func (c *comicService) GetComicByID(ctx context.Context, comicId int64) (*domain.Comic, error) {
	comic, err := c.cr.Get(ctx, comicId)
	if err != nil {
		return nil, err
	}

	cover, err := c.cvr.GetActive(ctx, comicId)
	if err != nil {
		return nil, err
	}

	comic.CoverUrl = cover.URL

	return comic, nil
}

func (c *comicService) NewComic(ctx context.Context, title, description, author, artist, comicStatus, comicType string, names []string) (*domain.Comic, error) {
	comic, err := domain.NewComic(title, description, author, artist, comicStatus, comicType)
	if err != nil {
		return nil, fmt.Errorf("%w: %s", ErrValidation, err.Error())
	}

	genres, err := domain.NewMassGenre(names)
	if err != nil {
		return nil, fmt.Errorf("%w: %s", ErrValidation, err.Error())
	}

	// Should these 2 be inside different goroutine?
	if err = c.cr.Create(ctx, comic); err != nil {
		return nil, err
	}

	// Upsert regardless of comic insert success
	if err = c.gr.MassUpsert(ctx, genres); err != nil {
		return nil, err
	}

	genreIds := make([]int64, 0)
	for _, genre := range genres {
		genreIds = append(genreIds, genre.ID)
	}

	// How do I maintain the transaction?
	if err = c.cgr.MassUpsert(ctx, comic.ID, genreIds); err != nil {
		return nil, err
	}

	comic.Genres = names

	return comic, nil
}
