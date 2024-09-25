package config

import (
  "log"
  "os"

  "github.com/joho/godotenv"
  "github.com/spf13/cast"
)

type Config struct {
  USER_SERVICE         string
  DB_HOST              string
  DB_PORT              string
  DB_USER              string
  DB_PASSWORD          string
  DB_NAME              string
  ACCESS_TOKEN_SECRET  string
  REFRESH_TOKEN_SECRET string
  SIGNING_KEY          string
}

func LoadConfig() Config {
  if err := godotenv.Load(".env"); err != nil {
    log.Println("error loading .env file or not found", err)
  }

  config := Config{}

  config.USER_SERVICE = cast.ToString(coalesce("USER_SERVICE", ":50053"))
  config.DB_HOST = cast.ToString(coalesce("DB_HOST", "localhost"))
  config.DB_PORT = cast.ToString(coalesce("DB_PORT", "5432"))
  config.DB_USER = cast.ToString(coalesce("DB_USER", "macbookpro"))
  config.DB_PASSWORD = cast.ToString(coalesce("DB_PASSWORD", "1111"))
  config.DB_NAME = cast.ToString(coalesce("DB_NAME", "testuzb_question_service"))
  config.SIGNING_KEY = cast.ToString(coalesce("SIGNING_KEY", "secret"))

  return config
}

func coalesce(key string, defaultValue interface{}) interface{} {
  value, exists := os.LookupEnv(key)
  if exists {
    return value
  }
  return defaultValue
}