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

func PrepareRequest(c *gin.Context, method int) (models.RequestHost, error) {
	var url string
	var path string
	var paylod *bytes.Buffer
	var requestMethod string

	s := c.Param("server")
	switch s {
	case "server1":
		if method == models.GET {
			requestMethod = "GET"
			url = config.SERVER1_DEFAULT_HOST
			path = config.SERVER1_PATH
			id := strings.Replace(c.Param("id"), "/", "", 1)
			if id != "" {
				path = fmt.Sprintf("%s?userId=%s", path, id)
			} else {
				return models.RequestHost{}, fmt.Errorf("err")
			}
		} else if method == models.POST {
			requestMethod = "POST"
			paylod = ClientBodyHandler(c)
			// TODO: Montar url...
		}
	case "server2":
		if method == models.GET {
			requestMethod = "GET"
			url = config.SERVER2_DEFAULT_HOST
			path = config.SERVER2_PATH
			id, _ := strconv.Atoi(strings.Replace(c.Param("id"), "/", "", 1))
			if id > 0 {
				path = fmt.Sprintf("%s?book_id=%d", path, id)
			} else {
				path = fmt.Sprintf("%s/all", path)
			}
		} else if method == models.POST {
			requestMethod = "POST"
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

// ? Deve fazer parte do contexto?
// ? Precisa fazer parte do contexto?
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