package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/spf13/cast"
)

type Config struct {
	HTTP_PORT         string
	ACCESS_TOKEN      string
	USER_SERVICE_PORT string
	CONTENT_SERVICE_PORT string
}

func Load() Config {
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found")
	}

	config := Config{}

	config.HTTP_PORT = cast.ToString(coalesce("HTTP_PORT", "passwrod"))
	config.CONTENT_SERVICE_PORT = cast.ToString(coalesce("CONTENT_SERVICE_PORT", "50051"))
	config.ACCESS_TOKEN = cast.ToString(coalesce("ACCESS_TOKEN", "postgres"))
	config.USER_SERVICE_PORT = cast.ToString(coalesce("USER_SERVICE_PORT", "50050"))
	return config
}

func coalesce(env string, defaultValue interface{}) interface{} {
	value, exists := os.LookupEnv(env)
	if !exists {
		return defaultValue
	}
	return value
}
