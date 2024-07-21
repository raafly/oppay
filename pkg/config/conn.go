package config

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/spf13/viper"
	_ "github.com/lib/pq"
)

func NewConfig() *sql.DB {
	vp := viper.New()
	vp.SetConfigFile(".env")
	vp.AddConfigPath("/home/saturna/Desktop/Developments/GOLANG/src/ewallet/")

	err := vp.ReadInConfig()
	if err != nil {
		log.Fatalf("failed to read %v", err)
	}

	host := vp.GetString("DB_HOST")
	port := vp.GetString("DB_PORT")
	username := vp.GetString("DB_USER")
	password := vp.GetString("DB_PASSWORD")
	dbname := vp.GetString("DB_NAME")

	connStr := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", host, port, username, password, dbname)
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatalf("failed to connect db %v", err)
	}

	return db
}