package server1

import (
	"Api-Gateway-lcs42/routers/server1/handlersServer1"

	"github.com/gin-gonic/gin"
)

func Server1(rg *gin.RouterGroup) {
	server1 := rg.Group("/")
	server1.GET("/health", handlersServer1.HealthCheck)
	server1.GET("/createUser", handlersServer1.CreateUser)
}
