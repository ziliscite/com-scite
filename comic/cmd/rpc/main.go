package main

import "github.com/ziliscite/com-scite/comic/pkg/db"

func main() {
	cfg := getConfig()
	db.AutoMigrate(cfg.db.dsn)
}
