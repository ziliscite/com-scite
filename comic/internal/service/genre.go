package service

import (
	"context"
	"fmt"
	"github.com/ziliscite/com-scite/comic/internal/domain"
	"github.com/ziliscite/com-scite/comic/internal/repository"
)

type GenreService interface {
	GetAllGenres(ctx context.Context) ([]domain.Genre, error)
	UpsertGenres(ctx context.Context, names []string) ([]*domain.Genre, error)
	InsertGenre(ctx context.Context, name string) (*domain.Genre, error)
	Edit(ctx context.Context, genreId int64, newName string) (*domain.Genre, error)
	Delete(ctx context.Context, genreId int64) error
}

type genreService struct {
	gr repository.GenreRepository
}

func (g *genreService) GetAllGenres(ctx context.Context) ([]domain.Genre, error) {
	return g.gr.GetAll(ctx)
}

func (g *genreService) UpsertGenres(ctx context.Context, names []string) ([]*domain.Genre, error) {
	genres, err := domain.NewMassGenre(names)
	if err != nil {
		return nil, fmt.Errorf("%w: %s", ErrValidation, err.Error())
	}

	if err = g.gr.MassUpsert(ctx, genres); err != nil {
		return nil, err
	}

	return genres, nil
}

func (g *genreService) InsertGenre(ctx context.Context, name string) (*domain.Genre, error) {
	genre, err := domain.NewGenre(name)
	if err != nil {
		return nil, fmt.Errorf("%w: %s", ErrValidation, err.Error())
	}

	if err = g.gr.New(ctx, genre); err != nil {
		return nil, err
	}

	return genre, nil
}

func (g *genreService) Edit(ctx context.Context, genreId int64, newName string) (*domain.Genre, error) {
	if genreId == 0 {
		return nil, fmt.Errorf("%w: genre id cannot be empty", ErrValidation)
	}

	genre, err := domain.NewGenre(newName)
	if err != nil {
		return nil, fmt.Errorf("%w: %s", ErrValidation, err.Error())
	}

	genre.ID = genreId

	if err = g.gr.Edit(ctx, genre); err != nil {
		return nil, err
	}

	return genre, nil
}

func (g *genreService) Delete(ctx context.Context, genreId int64) error {
	if genreId == 0 {
		return fmt.Errorf("%w: genre id cannot be empty", ErrValidation)
	}

	return g.gr.Delete(ctx, genreId)
}
