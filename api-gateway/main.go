package main

import (
	"Api-Gateway-lcs42/config"
	"Api-Gateway-lcs42/routers"
)

func init() {
	config.LoadConfig()
}

func main() {
	routers.Run()
}
