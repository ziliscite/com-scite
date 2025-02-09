package main

import (
	amqp "github.com/rabbitmq/amqp091-go"
	"github.com/ziliscite/micro-auth/mailer/internal"
	"log/slog"
	"os"
)

func main() {
	cfg := getConfig()

	conn, err := amqp.Dial("amqp://guest:guest@rabbitmq")
	if err != nil {
		slog.Error(err.Error())
		os.Exit(1)
	}
	defer conn.Close()

	mailer := internal.New(cfg.mail.host, cfg.mail.port, cfg.mail.username, cfg.mail.password, cfg.mail.sender)

	msrv, err := newService(conn, mailer)
	if err != nil {
		slog.Error(err.Error())
		os.Exit(1)
	}

	err = msrv.listen()
	if err != nil {
		slog.Error(err.Error())
		os.Exit(1)
	}
}
