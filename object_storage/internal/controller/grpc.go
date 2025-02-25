package controller

import (
	"bytes"
	"context"
	"errors"
	"fmt"
	"github.com/ziliscite/com-scite/object_storage/internal/repository"
	pb "github.com/ziliscite/com-scite/object_storage/pkg/protobuf"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"io"
	"log/slog"
	"net"
	"os"
)

// 10MB
const maxImageSize = 10 << 20

type GrpcServer struct {
	pb.UnimplementedUploadServiceServer
	st repository.Write
}

func NewGrpcServer(st repository.Write) *GrpcServer {
	return &GrpcServer{
		st: st,
	}
}

func (s *GrpcServer) UploadImage(stream pb.UploadService_UploadImageServer) error {
	// receive the first request, which contains the metadata
	req, err := stream.Recv()
	if err != nil {
		slog.Error("cannot receive image info", "error", err.Error())
		return status.Errorf(codes.Unknown, "cannot receive image info")
	}

	fn := req.GetMetadata().GetFilename()
	ty := req.GetMetadata().GetTypes()

	// create a new byte buffer to store chunks
	imageData := bytes.Buffer{}
	imageSize := 0

	for {
		req, err = stream.Recv()
		if errors.Is(err, io.EOF) {
			slog.Info("no more data")
			break
		}
		if err != nil {
			slog.Error("cannot receive image info", "error", err.Error())
			return status.Errorf(codes.Unknown, "cannot receive image info")
		}

		chunk := req.GetChunk()
		size := len(chunk)

		slog.Info("received a chunk", "size", size)

		imageSize += size
		if imageSize > maxImageSize {
			return status.Errorf(codes.InvalidArgument, "image is too large: %d > %d", imageSize, maxImageSize)
		}

		// appends the chunk to the image buffer
		if _, err = imageData.Write(chunk); err != nil {
			return status.Errorf(codes.Internal, "cannot write chunk data: %v", err)
		}
	}

	signedUrl, err := s.st.Save(fn, ty, imageData)
	if err != nil {
		return status.Errorf(codes.Internal, "cannot save image to the store: %v", err)
	}

	// send the response to client
	if err = stream.SendAndClose(&pb.UploadImageResponse{
		SignedUrl: signedUrl,
	}); err != nil {
		return status.Errorf(codes.Unknown, "cannot send response: %v", err)
	}

	slog.Info("saved image", "url", signedUrl, "size", imageSize)
	return nil
}

func (s *GrpcServer) DeleteImage(ctx context.Context, req *pb.DeleteImageRequest) (*pb.Nothing, error) {
	if err := s.st.Delete(req.SignedUrl); err != nil {
		slog.Error("delete image failed", "error", err.Error())
		return nil, status.Errorf(codes.Internal, "cannot delete image from the store: %v", err)
	}

	return &pb.Nothing{}, nil
}

func (s *GrpcServer) Run(port int) {
	listen, err := net.Listen("tcp", fmt.Sprintf("0.0.0.0:%v", port))
	if err != nil {
		slog.Error(err.Error())
		os.Exit(1)
	}
	defer listen.Close()

	srv := grpc.NewServer()

	pb.RegisterUploadServiceServer(srv, s)

	if err = srv.Serve(listen); err != nil {
		slog.Error(err.Error())
		os.Exit(1)
	}
}
