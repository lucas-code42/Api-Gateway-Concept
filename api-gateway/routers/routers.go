package routers

import (
	"Api-Gateway-lcs42/config"
	"Api-Gateway-lcs42/routers/server1"
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
)

var httpEngine = gin.Default()

// Run initialize api
func Run() {
	mountRoutes()

	err := httpEngine.Run(fmt.Sprintf(":%s", config.PORT))
	if err != nil {
		log.Fatal("Could not start APIGATEWAY")
	}
}

func mountRoutes() {
	SERVER_1_ROUTE_GROUP := httpEngine.Group(fmt.Sprintf("%s/server1", config.SERVER_DEFAULT_PATH))
	server1.RouteServer1(SERVER_1_ROUTE_GROUP)
}
