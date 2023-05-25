package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

var (
	PORT         = ""
	DEFAULT_PATH = ""
	JWT_KEY      = ""
	JWT_ISSUER   = ""
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
}
