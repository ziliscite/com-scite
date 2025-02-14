package service

import (
	"context"
	"github.com/ziliscite/com-scite/comic/internal/domain"
	"github.com/ziliscite/com-scite/comic/internal/repository"
)

type ComicService interface {
	GetComicByID(ctx context.Context, comicId int64) (*domain.Comic, error)
}

type comicService struct {
	cr repository.ComicRepository
}

func (c *comicService) GetComicByID(ctx context.Context, comicId int64) (*domain.Comic, error) {
	return c.cr.Get(ctx, comicId)
}
