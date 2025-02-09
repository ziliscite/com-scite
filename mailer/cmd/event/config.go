package main

import (
	"flag"
	"os"
	"sync"
)

type Config struct {
	mail struct {
		host     string
		port     int
		username string
		password string
		sender   string
	}
}

var (
	instance Config
	once     sync.Once
)

func getConfig() Config {
	once.Do(func() {
		instance = Config{}

		flag.StringVar(&instance.mail.host, "smtp-host", "mailhog", "SMTP host")
		flag.IntVar(&instance.mail.port, "smtp-port", 1025, "SMTP port")
		flag.StringVar(&instance.mail.username, "smtp-username", os.Getenv("SMTP_USERNAME"), "SMTP username")
		flag.StringVar(&instance.mail.password, "smtp-password", os.Getenv("SMTP_PASSWORD"), "SMTP password")
		flag.StringVar(&instance.mail.sender, "smtp-sender", os.Getenv("SMTP_SENDER"), "SMTP sender")

		flag.Parse()
	})

	return instance
}
