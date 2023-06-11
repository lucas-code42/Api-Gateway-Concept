package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

var (
	PORT                 = ""
	SERVER_DEFAULT_PATH  = ""
	APIGATEWAY_JWT_KEY   = ""
	JWT_ISSUER_SERVER_1  = ""
	DEFAULT_HOST_SERVER1 = ""
)

// LoadConfig loads the .env file
func LoadConfig() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	PORT = os.Getenv("PORT")
	SERVER_DEFAULT_PATH = os.Getenv("SERVER_DEFAULT_PATH")
	APIGATEWAY_JWT_KEY = os.Getenv("APIGATEWAY_JWT_KEY")
	JWT_ISSUER_SERVER_1 = os.Getenv("JWT_ISSUER_SERVER_1")
	DEFAULT_HOST_SERVER1 = os.Getenv("DEFAULT_HOST_SERVER1")
}