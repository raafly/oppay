package db

import (
	"fmt"
	"log"
	"time"

	"github.com/raafly/ewallet/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func NewPostgres(cfg *config.App) *gorm.DB {
	name := cfg.Postgres.Name
	user := cfg.Postgres.User
	pass := cfg.Postgres.Pass
	host := cfg.Postgres.Host
	port := cfg.Postgres.Port

	dsn := fmt.Sprintf("host=%s user=%s dbname=%s port=%s password=%s", host, user, name, port, pass)
	log.Println(dsn)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Printf("failed init db %s", err)
	}

	sql, _ := db.DB()
	sql.SetConnMaxIdleTime(5 *time.Minute)
	sql.SetConnMaxLifetime(60 *time.Minute)
	sql.SetMaxIdleConns(10)
	sql.SetMaxOpenConns(100)

	return db
}