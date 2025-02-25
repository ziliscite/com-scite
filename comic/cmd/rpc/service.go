package main

import (
	"fmt"
	"log/slog"
	"net"
	"os"

	"github.com/ziliscite/com-scite/comic/internal/controller"

	pb "github.com/ziliscite/com-scite/comic/pkg/protobuf"

	"google.golang.org/grpc"
)

type ComicServer struct {
	cc  *controller.ComicController
	cvc *controller.CoverController
}

func NewComicServer(cc *controller.ComicController, cvc *controller.CoverController) *ComicServer {
	return &ComicServer{
		cc:  cc,
		cvc: cvc,
	}
}

func (cs *ComicServer) Serve(port int) error {
	listen, err := net.Listen("tcp", fmt.Sprintf("0.0.0.0:%v", port))
	if err != nil {
		slog.Error(err.Error())
		os.Exit(1)
	}
	defer listen.Close()

	srv := grpc.NewServer()

	pb.RegisterComicServiceServer(srv, cs.cc)
	pb.RegisterCoverServiceServer(srv, cs.cvc)

	return srv.Serve(listen)
}
