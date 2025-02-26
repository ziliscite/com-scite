package controller

import (
	"context"
	"errors"
	"log/slog"

	"github.com/ziliscite/com-scite/comic/internal/repository"
	"github.com/ziliscite/com-scite/comic/internal/service"

	pb "github.com/ziliscite/com-scite/comic/pkg/protobuf"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type ComicController struct {
	pb.UnimplementedComicServiceServer
	cs  service.ComicService
	cvs service.CoverService
}

func NewComicController(cs service.ComicService, cvs service.CoverService) *ComicController {
	return &ComicController{
		cs: cs, cvs: cvs,
	}
}

func (cc *ComicController) InsertComic(ctx context.Context, req *pb.InsertComicRequest) (*pb.InsertComicResponse, error) {
	comic, err := cc.cs.NewComic(
		ctx, req.GetTitle(), req.GetDescription(), req.GetAuthor(),
		req.GetArtist(), req.GetStatus(), req.GetType(), req.GetGenres(),
	)
	if err != nil {
		slog.Error("NewComic failed", "error", err.Error())
		switch {
		case errors.Is(err, context.DeadlineExceeded):
			return nil, status.Error(codes.DeadlineExceeded, err.Error())
		case errors.Is(err, context.Canceled):
			return nil, status.Error(codes.Canceled, err.Error())
		case errors.Is(err, service.ErrValidation):
			return nil, status.Error(codes.InvalidArgument, err.Error())
		case errors.Is(err, repository.ErrDuplicate):
			return nil, status.Error(codes.AlreadyExists, err.Error())
		default:
			return nil, status.Error(codes.Internal, err.Error())
		}
	}

	return &pb.InsertComicResponse{
		Comic: &pb.Comic{
			Id:          comic.ID,
			Title:       comic.Title,
			Slug:        comic.Slug,
			Description: comic.Description,
			Author:      comic.Author,
			Artist:      comic.Artist,
			Status:      comic.Status.String(),
			Type:        comic.Type.String(),
			Genres:      comic.Genres,
		},
	}, nil
}

func (cc *ComicController) GetComicBySlug(ctx context.Context, req *pb.GetComicBySlugRequest) (*pb.GetComicBySlugResponse, error) {
	comic, err := cc.cs.GetComicBySlug(ctx, req.GetSlug())
	if err != nil {
		slog.Error("Get Comic failed", "error", err.Error())
		switch {
		case errors.Is(err, context.DeadlineExceeded):
			return nil, status.Error(codes.DeadlineExceeded, err.Error())
		case errors.Is(err, context.Canceled):
			return nil, status.Error(codes.Canceled, err.Error())
		case errors.Is(err, repository.ErrNotFound):
			return nil, status.Error(codes.NotFound, err.Error())
		default:
			return nil, status.Error(codes.Internal, err.Error())
		}
	}

	// get cover as well
	cover, err := cc.cvs.GetActive(ctx, comic.ID)
	if err != nil {
		switch {
		case errors.Is(err, service.ErrNotFound):
			return nil, status.Error(codes.NotFound, err.Error())
		default:
			return nil, status.Error(codes.Internal, err.Error())
		}
	}

	return &pb.GetComicBySlugResponse{
		Comic: &pb.Comic{
			Id:          comic.ID,
			Title:       comic.Title,
			Slug:        comic.Slug,
			Description: comic.Description,
			Author:      comic.Author,
			Artist:      comic.Artist,
			Status:      comic.Status.String(),
			Type:        comic.Type.String(),
			Genres:      comic.Genres,
		},
		Cover: &pb.Cover{
			Id:        cover.ID,
			ComicId:   cover.ComicID,
			FileKey:   cover.FileKey,
			IsCurrent: cover.IsCurrent,
		},
	}, nil
}
