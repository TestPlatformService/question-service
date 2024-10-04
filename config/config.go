package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/spf13/cast"
)

type Config struct {
	QUESTION_SERVICE string
	DB_HOST      string
	DB_PORT      string
	DB_USER      string
	DB_PASSWORD  string
	DB_NAME      string
	MDB_ADDRESS  string
	MDB_NAME     string
	USER_SERVICE string
}

func LoadConfig() Config {
	if err := godotenv.Load(".env"); err != nil {
		log.Println("error loading .env file or not found", err)
	}

	config := Config{}

	config.QUESTION_SERVICE = cast.ToString(coalesce("QUESTION_SERVICE", ":50053"))
	config.DB_HOST = cast.ToString(coalesce("PDB_HOST", "localhost"))
	config.DB_PORT = cast.ToString(coalesce("PDB_PORT", "5432"))
	config.DB_USER = cast.ToString(coalesce("PDB_USER", "postgres"))
	config.DB_PASSWORD = cast.ToString(coalesce("PDB_PASSWORD", "1111"))
	config.DB_NAME = cast.ToString(coalesce("PDB_NAME", "testuzb1_question_service"))
	config.MDB_ADDRESS = cast.ToString(coalesce("MDB_ADDRESS", "mongodb://localhost:27017"))
	config.MDB_NAME = cast.ToString(coalesce("MDB_NAME", "testuzb_question_service"))
	config.USER_SERVICE = cast.ToString(coalesce("USER_SERVICE", ""))

	return config
}

func coalesce(key string, defaultValue interface{}) interface{} {
	value, exists := os.LookupEnv(key)
	if exists {
		return value
	}
	return defaultValue
}
