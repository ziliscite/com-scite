package main

import (
	"context"
	"errors"
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
	authClient, err := grpc.NewClient("auth:50051", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		slog.Error("Failed to connect to token service client", "error", err)
		os.Exit(1)
	}
	defer authClient.Close()

	activationClient, err := grpc.NewClient("activation:50051", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		slog.Error("Failed to connect to activation service client", "error", err)
		os.Exit(1)
	}
	defer activationClient.Close()

	comicClient, err := grpc.NewClient("comic:50051", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		slog.Error("Failed to connect to activation service client", "error", err)
		os.Exit(1)
	}
	defer comicClient.Close()

	authServ := pb.NewAuthServiceClient(authClient)
	activationServ := pb.NewActivationServiceClient(activationClient)
	comicServ := pb.NewComicServiceClient(comicClient)

	app := &applications{
		auc: authServ,
		atc: activationServ,
		cc:  comicServ,
	}

	server := &http.Server{Addr: ":80", Handler: app.routes()}

	done := make(chan os.Signal, 1)
	signal.Notify(done, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		if err := server.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			panic(err)
		}
	}()

	<-done

	if err = server.Shutdown(context.Background()); err != nil {
		slog.Error(err.Error())
		os.Exit(1)
	}
}
