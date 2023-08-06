package handlers

import (
	"Api-Gateway-lcs42/config"
	"fmt"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
)

const (
	GET    = 1
	POST   = 2
	PUT    = 3
	DELETE = 4
)

type Host struct {
	Url  string
	Path string
}

func PrepareRequest(c *gin.Context, method int) (Host, error) {
	var url string
	var path string

	s := c.Param("server")
	switch s {
	case "server1":
		if method == GET {
			url = config.SERVER1_DEFAULT_HOST
			path = config.SERVER1_PATH
			id := strings.Replace(c.Param("id"), "/", "", 1)
			if id != "" {
				path = fmt.Sprintf("%s?userId=%s", path, id)
			} else {
				return Host{}, fmt.Errorf("err")
			}
		}
	case "server2":
		if method == GET {
			url = config.SERVER2_DEFAULT_HOST
			path = config.SERVER2_PATH
			id, _ := strconv.Atoi(strings.Replace(c.Param("id"), "/", "", 1))
			if id > 0 {
				path = fmt.Sprintf("%s?book_id=%d", path, id)
			} else {
				path = fmt.Sprintf("%s/all", path)
			}
		}
	default:
		return Host{}, fmt.Errorf("err")
	}

	return Host{url, path}, nil
}

func BodyHandler() {

}
