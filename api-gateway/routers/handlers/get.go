package handlers

import (
	"Api-Gateway-lcs42/config"
	"Api-Gateway-lcs42/routers/tools"
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
)

type host struct {
	url  string
	path string
}

func ServerInterfaceGet(c *gin.Context) {
	start := time.Now()

	serverHost, err := mountHost(c)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "internal error"})
	}

	r, err := tools.GetRequest(serverHost.url, serverHost.path, fmt.Sprintf("%v", c.Keys["jwt"]))
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "internal error"})
	}

	r.ExecutionTime = time.Duration(time.Since(start).Milliseconds())
	c.JSON(http.StatusOK, gin.H{"data": r})
}

func mountHost(c *gin.Context) (host, error) {
	var url string
	var path string

	s := c.Param("server")
	switch s {
	case "server1":
		url = config.SERVER1_DEFAULT_HOST
		path = config.SERVER1_PATH
		id := strings.Replace(c.Param("id"), "/", "", 1)
		if id != "" {
			path = fmt.Sprintf("%s?userId=%s", path, id)
		} else {
			return host{}, fmt.Errorf("err")
		}
	case "server2":
		url = config.SERVER2_DEFAULT_HOST
		path = config.SERVER2_PATH
		id, _ := strconv.Atoi(strings.Replace(c.Param("id"), "/", "", 1))
		if id > 0 {
			path = fmt.Sprintf("%s?book_id=%d", path, id)
		} else {
			path = fmt.Sprintf("%s/all", path)
		}
	default:
		return host{}, fmt.Errorf("err")
	}

	return host{url, path}, nil
}
