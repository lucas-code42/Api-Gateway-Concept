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
	httpEngine := gin.New()
	httpEngine.GET(fmt.Sprintf("%s/health", config.DEFAULT_PATH), handlers.HealthCheck)
	httpEngine.GET(fmt.Sprintf("%s/authenticate", config.DEFAULT_PATH), handlers.DeliveryToken)
	httpEngine.POST(fmt.Sprintf("%s/checkToken", config.DEFAULT_PATH), handlers.VerifyToken)
	// apiGateWayHandShake := httpEngine.Group("/")

	err := httpEngine.Run(fmt.Sprintf(":%s", config.PORT))
	if err != nil {
		log.Fatal("Could not start API")
	}
}
