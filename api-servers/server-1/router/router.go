package router

import (
	"fmt"
	"log"

	"github.com/api-server/lcs42/config"
	"github.com/api-server/lcs42/handlers"
	"github.com/gin-gonic/gin"
)

// StartApiEngine initialize api
func StartApiEngine() {
	// gin.SetMode(gin.DebugMode)

	// f, _ := os.Create("gin.log")
	// gin.DefaultWriter = io.MultiWriter(f)

	httpEngine := gin.Default()
	httpEngine.GET(fmt.Sprintf("%s/health", config.DEFAULT_PATH), handlers.HealthCheck)
	httpEngine.GET(fmt.Sprintf("%s/authenticate", config.DEFAULT_PATH), handlers.DeliveryToken)
	httpEngine.POST(fmt.Sprintf("%s/create", config.DEFAULT_PATH), handlers.CreateUser)
	
	// apiGateWayHandShake := httpEngine.Group("/")

	err := httpEngine.Run(fmt.Sprintf(":%s", config.PORT))
	if err != nil {
		log.Fatal("Could not start API")
	}
}
