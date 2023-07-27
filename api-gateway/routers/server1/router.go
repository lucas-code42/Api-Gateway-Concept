package server1

import (
	"Api-Gateway-lcs42/routers/server1/handlers"

	"github.com/gin-gonic/gin"
)

func Server1(rg *gin.RouterGroup) {
	server1 := rg.Group("/")
	server1.GET("/health", handlers.HealthCheck)
	server1.GET("/createUser", handlers.CreateUser)
}
