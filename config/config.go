package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	DB_HOST          string
	DB_PORT          string
	DB_USER          string
	DB_PASSWORD      string
	DB_NAME          string
	QUESTION_SERVICE string
}

func LoadConfig() Config {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	Config := Config{}
	Config.DB_HOST = os.Getenv("DB_HOST")
	Config.DB_PORT = os.Getenv("DB_PORT")
	Config.DB_USER = os.Getenv("DB_USER")
	Config.DB_PASSWORD = os.Getenv("DB_PASSWORD")
	Config.DB_NAME = os.Getenv("DB_NAME")
	Config.QUESTION_SERVICE = os.Getenv("QUESTION_SERVICE")
	return Config
}
