package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func Env(key string) string {
	// load .env file
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Error loading .env file")
	}
	return os.Getenv(key)
}

var (
	PORT       = Env("PORT")
	MONGO_URI  = Env("MONGO_URI")
	DATABASE   = Env("DATABASE")
	JWT_SECRET = Env("JWT_SECRET")
)
