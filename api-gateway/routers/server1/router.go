package server1

import (
	"Api-Gateway-lcs42/routers/server1/handlersServer1"

	"github.com/gin-gonic/gin"
)

func RouteServer1(rg *gin.RouterGroup) {
	s1 := rg.Group("/health")
	s1.GET("/", handlersServer1.HealthCheck)

}
