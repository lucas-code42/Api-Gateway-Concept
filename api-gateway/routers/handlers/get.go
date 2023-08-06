package handlers

import (
	"Api-Gateway-lcs42/routers/tools"
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func ServerInterfaceGet(c *gin.Context) {
	start := time.Now()

	serverHost, err := PrepareRequest(c, GET)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "internal error"})
		return
	}

	r, err := tools.DoRequest(
		serverHost.Url, serverHost.Path, fmt.Sprintf("%v", c.Keys["jwt"]), "GET", nil,
	)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "internal error"})
		return
	}

	r.ExecutionTime = time.Duration(time.Since(start).Milliseconds())
	c.JSON(http.StatusOK, gin.H{"data": r})
}
