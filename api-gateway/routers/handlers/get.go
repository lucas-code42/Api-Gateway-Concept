package handlers

import (
	"Api-Gateway-lcs42/config"
	"Api-Gateway-lcs42/routers/tools"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func ServerInterfaceGet(c *gin.Context) {
	start := time.Now()

	server := c.Param("serverName")

	var url string
	var path string

	switch server {
	case "server1":
		url = config.DEFAULT_HOST_SERVER1
		path = config.DEFAULT_HOST_SERVER1
	case "server2":
		url = config.DEFAULT_HOST_SERVER2
		path = "book" // * temp, should be in .env
	default:
		c.JSON(http.StatusNotFound, gin.H{"error": "unknown server"})
	}

	jwtToken, err := tools.GetJwt(server)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "cannot auth with server"})
	}

	r, err := tools.GetRequest(url, path, jwtToken.Token)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "internal error"})
	}
	r.ExecutionTime = time.Duration(time.Since(start).Milliseconds())

	c.JSON(http.StatusOK, gin.H{"data": r})
}
