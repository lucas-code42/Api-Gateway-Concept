package handlers

import (
	"Api-Gateway-lcs42/config"
	"Api-Gateway-lcs42/routers/server2/tools"
	"bytes"
	"encoding/json"

	"github.com/gin-gonic/gin"
)

func CreateBook(c *gin.Context) {
	var clientJson interface{}
	if err := c.BindJSON(&clientJson); err != nil {
		c.JSON(400, gin.H{"err": "unprocessable entity"})
		return
	}

	clientData := new(bytes.Buffer)
	json.NewEncoder(clientData).Encode(clientJson)

	r, err := tools.RequestServer2("POST", config.SERVER1_PATH, clientData)
	if err != nil {
		c.JSON(500, gin.H{"err": "Internal server error"})
		return
	}
	c.JSON(201, gin.H{"data": r})
}
