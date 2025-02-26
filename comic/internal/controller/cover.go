package controller

import (
	"bytes"
	"errors"
	"github.com/ziliscite/com-scite/comic/internal/service"
	pb "github.com/ziliscite/com-scite/comic/pkg/protobuf"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"io"
)

type CoverController struct {
	pb.UnimplementedCoverServiceServer
	cvs service.CoverService
}

func NewCoverController(cvs service.CoverService) *CoverController {
	return &CoverController{cvs: cvs}
}

const maxImageSize = 10 << 20

func (c *CoverController) UploadCover(stream pb.CoverService_UploadCoverServer) error {
	ctx := stream.Context()

	req, err := stream.Recv()
	if err != nil {
		return status.Errorf(codes.Unknown, "cannot receive image info")
	}

	fn := req.GetMetadata().GetFilename()
	ci := req.GetMetadata().GetComicId()

	imageData := bytes.Buffer{}
	imageSize := 0

	for {
		req, err = stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			return status.Errorf(codes.Unknown, "cannot receive image info")
		}

		chunk := req.GetChunk()
		size := len(chunk)

		imageSize += size
		if imageSize > maxImageSize {
			return status.Errorf(codes.InvalidArgument, "image is too large: %d > %d", imageSize, maxImageSize)
		}

		if _, err = imageData.Write(chunk); err != nil {
			return status.Errorf(codes.Internal, "cannot write chunk data: %v", err)
		}
	}

	cover, err := c.cvs.UploadImage(ctx, imageData, fn, ci)
	if err != nil {
		switch {
		case errors.Is(err, service.ErrValidation):
			return status.Errorf(codes.InvalidArgument, "cannot upload image: %v", err)
		case errors.Is(err, service.ErrDuplicate):
			return status.Errorf(codes.AlreadyExists, "cannot upload image: %v", err)
		case errors.Is(err, service.ErrNotFound):
			return status.Errorf(codes.NotFound, "cannot upload image: %v", err)
		default:
			return status.Errorf(codes.Internal, "cannot upload image: %v", err)
		}
	}

	// send the response to client
	if err = stream.SendAndClose(&pb.UploadCoverResponse{
		Cover: &pb.Cover{
			Id:        cover.ID,
			ComicId:   cover.ComicID,
			FileKey:   cover.FileKey,
			IsCurrent: cover.IsCurrent,
		},
	}); err != nil {
		return status.Errorf(codes.Unknown, "cannot send response: %v", err)
	}

	return nil
}
