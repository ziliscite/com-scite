package main

import (
	"github.com/ziliscite/com-scite/object_storage/internal/controller"
	"github.com/ziliscite/com-scite/object_storage/internal/repository"
	"github.com/ziliscite/com-scite/object_storage/pkg/encryptor"
	"os"

	"log"
)

func main() {
	enc := encryptor.NewEncryptor(os.Getenv("KEY"))
	store := repository.NewStore(enc)

	httpServer := controller.NewHttpServer(store)
	grpcServer := controller.NewGrpcServer(store)

	go grpcServer.Run(50051)

	if err := httpServer.Run(80); err != nil {
		log.Fatal(err)
	}
}
