package main

import (
	"context"
	"fmt"
	amqp "github.com/rabbitmq/amqp091-go"
	"github.com/ziliscite/micro-auth/token/internal/repository"
	"github.com/ziliscite/micro-auth/token/internal/service"
	"github.com/ziliscite/micro-auth/token/pkg/db"
	pb "github.com/ziliscite/micro-auth/token/pkg/protobuf"
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

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	pool, err := db.Open(ctx, cfg.db.dsn)
	if err != nil {
		slog.Error(err.Error())
		os.Exit(1)
	}

	tokenRepository := repository.NewTokenRepository(pool)
	tokenService := service.NewTokenService(tokenRepository)

	conn, err := amqp.Dial("amqp://guest:guest@rabbitmq")
	if err != nil {
		slog.Error(err.Error())
		os.Exit(1)
	}
	defer conn.Close()

	mailPublisher, err := service.NewPublisher(conn)
	if err != nil {
		slog.Error(err.Error())
		os.Exit(1)
	}

	// target is a dockerized service
	authClient, err := grpc.NewClient("auth:50051", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		slog.Error("Failed to connect to auth service client", "error", err)
		os.Exit(1)
	}
	defer authClient.Close()

	tkn := newClient(tokenService, pb.NewAuthServiceClient(authClient), mailPublisher)

	listen, err := net.Listen("tcp", fmt.Sprintf("0.0.0.0:%v", cfg.port))
	if err != nil {
		slog.Error(err.Error())
		os.Exit(1)
	}
	defer listen.Close()

	gsrv := grpc.NewServer()

	pb.RegisterActivationServiceServer(gsrv, tkn)
	if err = gsrv.Serve(listen); err != nil {
		slog.Error(err.Error())
		os.Exit(1)
	}
}
