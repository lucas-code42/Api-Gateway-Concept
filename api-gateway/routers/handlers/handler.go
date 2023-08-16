package handlers

import (
	"Api-Gateway-lcs42/routers/handlers/requests"
	"Api-Gateway-lcs42/routers/tools"
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

// ServerInterfaceGet faz a requisição GET o server de destino
func RequestInterface(c *gin.Context) {
	start := time.Now()

	fmt.Println(c.Request.Method)

	serverHost, err := requests.PrepareRequest(c, c.Request.Method)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "internal error"})
		return
	}

	r, err := tools.DoRequest(serverHost)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "internal error"})
		return
	}

	r.ExecutionTime = time.Duration(time.Since(start).Milliseconds())
	c.JSON(http.StatusOK, gin.H{"data": r})
}
