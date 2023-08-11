package handlers

import (
	"Api-Gateway-lcs42/config"
	"bytes"
	"encoding/json"
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

type RequestHost struct {
	Url     string
	Path    string
	Token   any // TODO: encontrar um jeito de ser string
	Method  int
	Payload *bytes.Buffer
}

func PrepareRequest(c *gin.Context, method int, token any) (RequestHost, error) {
	var url string
	var path string
	var paylod *bytes.Buffer

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
				return RequestHost{}, fmt.Errorf("err")
			}
		} else if method == POST {
			paylod = ClientBodyHandler(c)
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
		} else if method == POST {
			paylod = ClientBodyHandler(c)
		}
	default:
		return RequestHost{}, fmt.Errorf("err")
	}

	return RequestHost{
		Url:     url,
		Path:    path,
		Token:   token,
		Method:  method,
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
