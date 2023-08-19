package routers

import (
	"Api-Gateway-lcs42/config"
	"Api-Gateway-lcs42/routers/handlers"
	"Api-Gateway-lcs42/routers/handlers/middleware"
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
)

var httpEngine = gin.Default()

// Run initialize api
func Run() {
	httpEngine.Use(middleware.DummyMiddleware())
	mountRoutes()
	err := httpEngine.Run(fmt.Sprintf(":%s", config.PORT))
	if err != nil {
		log.Fatal("Could not start APIGATEWAY")
	}
}

func mountRoutes() {
	httpEngine.GET("/:server/", func(c *gin.Context) {
		c.JSON(200, map[string]any{"hello": c.Keys["jwt"]})
	})

	httpEngine.GET(fmt.Sprintf("%s/:server/*id", config.SERVER_DEFAULT_PATH), handlers.RequestApiServers)
	httpEngine.POST(fmt.Sprintf("%s/:server/", config.SERVER_DEFAULT_PATH), handlers.RequestApiServers)
	httpEngine.PUT(fmt.Sprintf("%s/:server/", config.SERVER_DEFAULT_PATH), handlers.RequestApiServers)
	httpEngine.DELETE(fmt.Sprintf("%s/:server/", config.SERVER_DEFAULT_PATH), handlers.RequestApiServers)
}
