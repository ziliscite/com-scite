package service

import (
	"bufio"
	"context"
	"github.com/ziliscite/com-scite/comic/internal/domain"
	pb "github.com/ziliscite/com-scite/comic/pkg/protobuf"
	"io"
	"log/slog"
	"os"
	"path/filepath"
	"time"

	"github.com/ziliscite/com-scite/comic/internal/repository"
)

type CoverService interface {
}

type coverService struct {
	cvr repository.CoverRepository
	us  pb.UploadServiceClient
}

func NewCoverService(cvr repository.CoverRepository) CoverService {
	return &coverService{cvr: cvr}
}

func (c *coverService) GetActive(ctx context.Context, comicId int64) (*domain.Cover, error) {
	panic("implement me")
}

// , filename string, comicId int64
func (c *coverService) UploadImage(ctx context.Context) (*domain.Cover, error) {
	file, err := os.Open("./external/knight.jpeg")
	if err != nil {
		slog.Error("cannot open image file: ", "error", err.Error())
		return nil, err
	}
	defer file.Close()

	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	// get the stream
	stream, err := c.us.UploadImage(ctx)
	if err != nil {
		return nil, err
	}

	// build metadata to be sent first
	req := &pb.UploadImageRequest{
		Data: &pb.UploadImageRequest_Metadata{
			Metadata: &pb.Metadata{
				Filename: filepath.Base("knight.jpeg"),
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
	reader := bufio.NewReader(file)
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

	cover := domain.NewCover(1, res.SignedUrl)
	if err = c.cvr.New(ctx, &cover); err != nil {
		return nil, err
	}

	return &cover, nil
}
