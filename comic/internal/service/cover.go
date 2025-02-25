package service

import (
	"bufio"
	"bytes"
	"context"
	"github.com/ziliscite/com-scite/comic/internal/domain"
	"github.com/ziliscite/com-scite/comic/internal/repository"
	pb "github.com/ziliscite/com-scite/comic/pkg/protobuf"
	"io"
	"log/slog"
)

type CoverService interface {
	UploadImage(ctx context.Context, image bytes.Buffer, filename string, comicId int64) (*domain.Cover, error)
}

type coverService struct {
	cvr repository.CoverRepository
	us  pb.UploadServiceClient
}

func NewCoverService(cvr repository.CoverRepository, us pb.UploadServiceClient) CoverService {
	return &coverService{cvr: cvr, us: us}
}

func (c *coverService) GetActive(ctx context.Context, comicId int64) (*domain.Cover, error) {
	panic("implement me")
}

func (c *coverService) UploadImage(ctx context.Context, image bytes.Buffer, filename string, comicId int64) (*domain.Cover, error) {
	stream, err := c.us.UploadImage(ctx)
	if err != nil {
		slog.Error("cannot create stream: ", "error", err.Error())
		return nil, err
	}

	// build metadata to be sent first
	req := &pb.UploadImageRequest{
		Data: &pb.UploadImageRequest_Metadata{
			Metadata: &pb.Metadata{
				Filename: filename,
				Types:    "cover",
			},
		},
	}

	// send the first request to the server
	if err = stream.Send(req); err != nil {
		slog.Error("cannot read chunk to buffer: ", "error", err.Error(), "message", stream.RecvMsg(nil))
		return nil, err
	}

	// create buffer to send a file
	reader := bufio.NewReader(&image)
	buffer := make([]byte, 1024)

	for {
		// read the data to the buffer
		n, err := reader.Read(buffer)
		if err == io.EOF {
			break
		}
		if err != nil {
			slog.Error("cannot read chunk to buffer: ", "error", err.Error())
			return nil, err
		}

		req := &pb.UploadImageRequest{
			Data: &pb.UploadImageRequest_Chunk{
				Chunk: buffer[:n],
			},
		}

		// send chunk to the server
		if err = stream.Send(req); err != nil {
			slog.Error("cannot send chunk to server: ", "error", err.Error())
			return nil, err
		}
	}

	res, err := stream.CloseAndRecv()
	if err != nil {
		slog.Error("cannot receive response: ", "error", err.Error())
		return nil, err
	}

	cover := domain.NewCover(comicId, res.SignedUrl)
	oldKey, err := c.cvr.New(ctx, &cover)
	if err != nil {
		switch {
		// case errors.Is(err, repository.ErrDuplicate):
		// bikin error wrapper lagi
		default:
			return nil, err
		}
	}

	// delete old cover from the fileserver
	if _, err = c.us.DeleteImage(ctx, &pb.DeleteImageRequest{SignedUrl: oldKey}); err != nil {
		slog.Error("cannot delete old cover: ", "error", err.Error())
		return nil, err
	}

	return &cover, nil
}
