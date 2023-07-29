package server2

import (
	"github.com/gin-gonic/gin"
	"Api-Gateway-lcs42/routers/server2/handlers"
)

func Server2(rg *gin.RouterGroup) {
	server2 := rg.Group("/")
	server2.GET("/authentication", handlers.CreateBook)

}