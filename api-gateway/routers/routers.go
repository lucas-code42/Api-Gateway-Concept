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
	httpEngine.GET("/:n/", func(c *gin.Context) {
		resp := map[string]string{"hello": "Deu bom"}
		c.JSON(200, resp)
	})
	httpEngine.GET(fmt.Sprintf("%s/:serverName/", config.SERVER_DEFAULT_PATH), handlers.ServerInterfaceGet)
	httpEngine.POST(fmt.Sprintf("%s/:serverName/", config.SERVER_DEFAULT_PATH), handlers.ServerInterfacePost)
	httpEngine.PUT(fmt.Sprintf("%s/:serverName/", config.SERVER_DEFAULT_PATH), handlers.ServerInterfacePut)
	httpEngine.DELETE(fmt.Sprintf("%s/:serverName/", config.SERVER_DEFAULT_PATH), handlers.ServerInterfaceDelete)
}
