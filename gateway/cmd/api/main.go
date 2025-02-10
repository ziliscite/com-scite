package main

import (
	"context"
	"errors"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
	pb "github.com/ziliscite/micro-auth/gateway/pkg/protobuf"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	r := chi.NewRouter()

	r.Use(middleware.Recoverer)
	r.Use(RateLimit)
	r.Use(middleware.Logger)
	r.Use(middleware.Heartbeat("/healthz"))

	r.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: true,
		MaxAge:           300,
	}))

	authClient, err := grpc.NewClient("auth:50051", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		slog.Error("Failed to connect to token service client", "error", err)
		os.Exit(1)
	}
	defer authClient.Close()

	authServ := pb.NewAuthServiceClient(authClient)

	r.Post("/v0/register", func(w http.ResponseWriter, r *http.Request) {
		// Example handler for a gRPC endpoint
		var requestBody struct {
			Username string `json:"username"`
			Email    string `json:"email"`
			Password string `json:"password"`
		}

		err := readBody(w, r, &requestBody)
		if err != nil {
			sendError(w, http.StatusBadRequest, err)
			return
		}

		resp, err := authServ.Register(r.Context(), &pb.RegisterRequest{
			Username: requestBody.Username,
			Email:    requestBody.Email,
			Password: requestBody.Password,
		})
		if err != nil {
			sendError(w, http.StatusInternalServerError, err)
			return
		}

		err = writeJSON(w, http.StatusOK, resp)
		if err != nil {
			sendError(w, http.StatusInternalServerError, err)
			return
		}
	})

	server := &http.Server{
		Addr:    ":80",
		Handler: r,
	}

	// Graceful shutdown setup
	done := make(chan os.Signal, 1)
	signal.Notify(done, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		if err := server.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			panic(err)
		}
	}()

	<-done
	server.Shutdown(context.Background())
}
