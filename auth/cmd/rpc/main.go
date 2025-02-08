package main

import (
	"context"
	"github.com/ziliscite/micro-auth/auth/internal/repository"
	"github.com/ziliscite/micro-auth/auth/internal/service"
	"github.com/ziliscite/micro-auth/auth/pkg/db"
	"google.golang.org/grpc"
	"log/slog"
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
	}

	userRepository := repository.NewUserRepository(pool)
	userService := service.NewUserService(userRepository)

	auth := NewService(userService, cfg.jwtSecrets)

	server := grpc.NewServer()
}
