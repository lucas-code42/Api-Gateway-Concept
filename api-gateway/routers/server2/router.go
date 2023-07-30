package server2

import (
	"Api-Gateway-lcs42/routers/server2/handlers"

	"github.com/gin-gonic/gin"
)

func Server2(rg *gin.RouterGroup) {
	server2 := rg.Group("/")
	server2.GET("/authentication", handlers.CreateBook)

}
