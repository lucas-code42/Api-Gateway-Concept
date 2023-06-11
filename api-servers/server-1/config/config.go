package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

var (
	PORT           = ""
	DEFAULT_PATH   = ""
	JWT_KEY        = ""
	JWT_ISSUER     = ""
	REDIS_PASSWORD = ""
	REDIS_HOST     = ""
	REDIS_PORT     = ""
	SECURITY_KEY   = ""
)

// LoadConfig loads the .env file
func LoadConfig() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	PORT = os.Getenv("PORT")
	DEFAULT_PATH = os.Getenv("SERVER_DEFAULT_PATH")
	JWT_KEY = os.Getenv("JWT_KEY")
	JWT_ISSUER = os.Getenv("JWT_ISSUER")
	REDIS_PASSWORD = os.Getenv("REDIS_PASSWORD")
	REDIS_HOST = os.Getenv("REDIS_HOST")
	REDIS_PORT = os.Getenv("REDIS_PORT")
	SECURITY_KEY = os.Getenv("SECURITY_KEY")
}
