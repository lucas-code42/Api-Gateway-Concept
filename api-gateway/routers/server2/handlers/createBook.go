package handlers

import (
	"Api-Gateway-lcs42/routers/server2/tools"
	"fmt"

	"github.com/gin-gonic/gin"
)

func CreateBook(c *gin.Context) {
	fmt.Println(tools.GetAuthentication())
	c.JSON(200, gin.H{"data": true})
}
