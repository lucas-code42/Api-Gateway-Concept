package handlers

import (
	"Api-Gateway-lcs42/config"
	"Api-Gateway-lcs42/routers/tools"
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func ServerInterfaceGet(c *gin.Context) {
	start := time.Now()

	var url string
	var path string

	s := c.Param("server")
	switch s {
	case "server1":
		url = config.SERVER1_DEFAULT_HOST
		path = config.SERVER1_PATH
	case "server2":
		url = config.SERVER2_DEFAULT_HOST
		path = config.SERVER2_PATH
	default:
		c.JSON(http.StatusNotFound, gin.H{"error": "unknown server"})
	}

	r, err := tools.GetRequest(url, path, fmt.Sprintf("%v", c.Keys["jwt"]))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "internal error"})
	}

	r.ExecutionTime = time.Duration(time.Since(start).Milliseconds())
	c.JSON(http.StatusOK, gin.H{"data": r})
}
