package main

import (
	"flag"
	"os"
	"sync"
	"time"
)

type Config struct {
	port       int
	jwtSecrets string
	db         struct {
		dsn         string
		maxConns    int
		maxIdleTime time.Duration
	}
}

var (
	instance Config
	once     sync.Once
)

func getConfig() Config {
	once.Do(func() {
		instance = Config{}

		flag.IntVar(&instance.port, "port", 50051, "gRPC server port")

		flag.StringVar(&instance.jwtSecrets, "secrets", os.Getenv("JWT_SECRETS"), "256 bytes of secrets")

		flag.StringVar(&instance.db.dsn, "db-dsn", os.Getenv("DB_DSN"), "PostgreSQL DSN")
		flag.IntVar(&instance.db.maxConns, "db-max-open-conns", 25, "PostgreSQL max connections")
		flag.DurationVar(&instance.db.maxIdleTime, "db-max-idle-time", 15*time.Minute, "PostgreSQL max connection idle time")

		flag.Parse()
	})

	return instance
}
