package handlers

import (
	"Api-Gateway-lcs42/routers/handlers/requests"
	"Api-Gateway-lcs42/routers/httpHandler"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

// RequestApiServers faz a requisição GET o server de destino
func RequestApiServers(c *gin.Context) {
	start := time.Now()

	serverHost, err := requests.PrepareRequest(c, c.Request.Method)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "internal error"})
		return
	}

	res, err := httpHandler.Request(serverHost)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "internal error"})
		return
	}

	res.ExecutionTime = time.Duration(time.Since(start).Milliseconds())
	c.JSON(http.StatusOK, gin.H{"data": res})
}
