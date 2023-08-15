package handlers

import (
	"Api-Gateway-lcs42/models"
	"Api-Gateway-lcs42/routers/handlers/requests"
	"Api-Gateway-lcs42/routers/tools"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

// ServerInterfaceGet faz a requisição GET o server de destino
func ServerInterfaceGet(c *gin.Context) {
	start := time.Now()

	serverHost, err :=  requests.PrepareRequest(c, models.GET)
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
