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

	// health
	httpEngine.GET(fmt.Sprintf("%s/health", config.DEFAULT_PATH), handlers.HealthCheck)
	
	// auth
	httpEngine.GET(fmt.Sprintf("%s/authenticate", config.DEFAULT_PATH), handlers.DeliveryToken)
	
	// user
	httpEngine.POST(fmt.Sprintf("%s/create", config.DEFAULT_PATH), handlers.CreateUser)
	httpEngine.GET(fmt.Sprintf("%s/user", config.DEFAULT_PATH), handlers.GetUser)


	err := httpEngine.Run(fmt.Sprintf(":%s", config.PORT))
	if err != nil {
		log.Fatal("Could not start API")
	}
}
