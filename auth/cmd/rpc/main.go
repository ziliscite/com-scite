package main

import (
	"context"
	"fmt"
	"github.com/ziliscite/micro-auth/auth/internal/repository"
	"github.com/ziliscite/micro-auth/auth/internal/service"
	"github.com/ziliscite/micro-auth/auth/pkg/db"
	pb "github.com/ziliscite/micro-auth/auth/pkg/protobuf"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log/slog"
	"net"
	"os"
	"time"
)

func main() {
	cfg := getConfig()
	db.AutoMigrate(cfg.db.dsn)

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	pool, err := db.Open(ctx, cfg.db.dsn)
	if err != nil {
		slog.Error("Failed to connect to database", "error", err)
		os.Exit(1)
	}

	userRepository := repository.NewUserRepository(pool)
	userService := service.NewUserService(userRepository)

	// target is a dockerized service
	tokenClient, err := grpc.NewClient("token:50051", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		slog.Error("Failed to connect to token service client", "error", err)
		os.Exit(1)
	}
	defer tokenClient.Close()

	auth := NewService(userService, pb.NewActivationServiceClient(tokenClient), cfg.jwtSecrets)

	listen, err := net.Listen("tcp", fmt.Sprintf("0.0.0.0:%v", cfg.port))
	if err != nil {
		slog.Error(err.Error())
		os.Exit(1)
	}
	defer listen.Close()

	srv := grpc.NewServer()

	pb.RegisterAuthServiceServer(srv, auth)
	if err = srv.Serve(listen); err != nil {
		slog.Error(err.Error())
		os.Exit(1)
	}
}
