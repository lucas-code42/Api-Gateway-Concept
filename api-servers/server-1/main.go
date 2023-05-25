package main

import (
	"github.com/api-server/lcs42/config"
	"github.com/api-server/lcs42/router"
)

func init() {
	config.LoadConfig()
}

func main() {
	router.StartApiEngine()
}
