package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

var (
	// api gataway
	PORT                = ""
	SERVER_DEFAULT_PATH = ""
	APIGATEWAY_KEY      = ""
	// server1
	SERVER1_DEFAULT_HOST = ""
	SERVER1_AUTH_KEY     = ""
	SERVER1_PATH         = ""
	SERVER1_AUTH_PATH    = ""
	// server2
	SERVER2_DEFAULT_HOST = ""
	SERVER2_PATH         = ""
	SERVER2_AUTH_PATH    = ""
)

// LoadConfig loads the .env file
func LoadConfig() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// api gataway
	PORT = os.Getenv("PORT")
	SERVER_DEFAULT_PATH = os.Getenv("SERVER_DEFAULT_PATH")
	APIGATEWAY_KEY = os.Getenv("APIGATEWAY_KEY")
	// server1
	SERVER1_DEFAULT_HOST = os.Getenv("SERVER1_DEFAULT_HOST")
	SERVER1_AUTH_KEY = os.Getenv("SERVER1_AUTH_KEY")
	SERVER1_PATH = os.Getenv("SERVER1_PATH")
	SERVER1_AUTH_PATH = os.Getenv("SERVER1_AUTH_PATH")
	// server2
	SERVER2_DEFAULT_HOST = os.Getenv("SERVER2_DEFAULT_HOST")
	SERVER2_PATH = os.Getenv("SERVER2_PATH")
	SERVER2_AUTH_PATH = os.Getenv("SERVER2_AUTH_PATH")
}
