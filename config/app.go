package config

import (
	"log"

	"github.com/spf13/viper"
)

const configPath = "/home/melvinger/Development/GOLANG/src/ewallet/"

func NewConfig() *App {
	v := viper.New()
	v.SetConfigFile(".env")
	v.AddConfigPath(configPath)
	if err := v.ReadInConfig(); err != nil {
		log.Printf("failed init viper %s", err)
	}

	return &App{
		Postgres: Postgres{
			Name: v.GetString("DB_NAME"),
			Host: v.GetString("DB_HOST"),
			Port: v.GetString("DB_PORT"),
			User: v.GetString("DB_USER"),
			Pass: v.GetString("DB_PASS"),
		},
	}
}
