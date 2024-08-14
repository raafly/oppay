package main

import (	"log"

	"github.com/raafly/ewallet/config"
	"github.com/raafly/ewallet/db"
)

func main() {
	cfg := config.NewConfig()
	db := db.NewPostgres(cfg)
	log.Println(db)
}
