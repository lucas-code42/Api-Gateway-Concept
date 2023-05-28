package routers

import (
	"Api-Gateway-lcs42/config"
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
	server1 := httpEngine.Group(fmt.Sprintf("%s/server1", config.DEFAULT_PATH))
	sayHi(server1)

}
