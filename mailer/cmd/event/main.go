package main

import (
	amqp "github.com/rabbitmq/amqp091-go"
	"github.com/ziliscite/micro-auth/mailer/internal"
	"log/slog"
	"os"
)

type client struct {
	rc *amqp.Connection
	mr *internal.Mailer
}

func main() {
	cfg := getConfig()

	conn, err := amqp.Dial("amqp://guest:guest@rabbitmq")
	if err != nil {
		slog.Error(err.Error())
		os.Exit(1)
	}
	defer conn.Close()

	mailer := internal.New(cfg.mail.host, cfg.mail.port, cfg.mail.username, cfg.mail.password, cfg.mail.sender)

	cl := &client{conn, mailer}
	cl.sendEmail()
}
