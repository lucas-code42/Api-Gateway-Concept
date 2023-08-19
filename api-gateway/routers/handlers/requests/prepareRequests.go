package requests

import (
	"Api-Gateway-lcs42/config"
	"Api-Gateway-lcs42/models"
	"bytes"
	"encoding/json"
	"fmt"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
)

// ? talvez não precise de url e path... manter apenas url...

func PrepareRequest(c *gin.Context, method string) (models.RequestHost, error) {
	var url string
	var path string
	var paylod *bytes.Buffer
	var requestMethod string

	s := c.Param("server")
	switch s {
	case "server1":
		url = config.SERVER1_DEFAULT_HOST
		path = config.SERVER1_PATH
		if method == "GET" {
			requestMethod = "GET"
			id := strings.Replace(c.Param("id"), "/", "", 1)
			if id != "" {
				path = fmt.Sprintf("%s?userId=%s", path, id)
			} else {
				return models.RequestHost{}, fmt.Errorf("err")
			}
		} else if method == "POST" || method == "PUT" {
			requestMethod = method
			paylod = ClientBodyHandler(c)
		}
	case "server2":
		url = config.SERVER2_DEFAULT_HOST
		if method == "GET" {
			requestMethod = method
			id, _ := strconv.Atoi(strings.Replace(c.Param("id"), "/", "", 1))
			if id > 0 {
				path = fmt.Sprintf("book/?book_id=%d", id)
			} else {
				path = "book/all"
			}
		} else if method == "POST" || method == "PUT" {
			requestMethod = method
			paylod = ClientBodyHandler(c)
			url = fmt.Sprintf("%s/%s", config.SERVER2_DEFAULT_HOST, config.SERVER2_PATH)
		}
	default:
		return models.RequestHost{}, fmt.Errorf("err")
	}

	return models.RequestHost{
		Url:     url,
		Path:    path,
		Token:   fmt.Sprintf("%s", c.Keys["jwt"]),
		Method:  requestMethod,
		Payload: paylod,
	}, nil
}

func ClientBodyHandler(c *gin.Context) *bytes.Buffer {
	var clientJson interface{}
	if err := c.BindJSON(&clientJson); err != nil {
		c.JSON(400, gin.H{"err": "unprocessable entity"})
		return nil
	}
	clientData := new(bytes.Buffer)
	json.NewEncoder(clientData).Encode(clientJson)
	return clientData
}
