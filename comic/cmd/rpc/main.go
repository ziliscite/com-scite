package main

import (
	"context"
	"log/slog"
	"os"
	"time"
	
	"github.com/ziliscite/com-scite/comic/internal/controller"
	"github.com/ziliscite/com-scite/comic/internal/repository"
	"github.com/ziliscite/com-scite/comic/internal/service"
	"github.com/ziliscite/com-scite/comic/pkg/db"
)

func main() {
	cfg := getConfig()
	db.AutoMigrate(cfg.db.dsn)

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	pool, err := db.Open(ctx, cfg.db.dsn)
	if err != nil {
		slog.Error("Failed to connect to database", "error", err)
		os.Exit(1)
	}

	comicRepository := repository.NewComicRepository(pool)
	genreRepository := repository.NewGenreRepository(pool)
	comicGenreRepository := repository.NewComicGenreRepository(pool)
	coverRepository := repository.NewCoverRepository(pool)

	comicService := service.NewComicService(comicRepository, coverRepository, genreRepository, comicGenreRepository)

	comicController := controller.NewComicController(comicService)

	server := NewComicServer(comicController)

	if err = server.Serve(cfg.port); err != nil {
		slog.Error(err.Error())
		os.Exit(1)
	}
}
