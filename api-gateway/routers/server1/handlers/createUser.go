package handlers

import (
	"Api-Gateway-lcs42/config"
	"Api-Gateway-lcs42/routers/server1/tools"

	"bytes"
	"encoding/json"

	"github.com/gin-gonic/gin"
)

func CreateUser(ctx *gin.Context) {
	var clientJson interface{}
	if err := ctx.BindJSON(&clientJson); err != nil {
		ctx.JSON(500, gin.H{"err": "deu bosta"})
		return
	}
	clientData := new(bytes.Buffer)
	json.NewEncoder(clientData).Encode(clientJson)

	r, err := tools.RequestServer("POST", config.SERVER1_CREATE_PATH, clientData)
	if err != nil {
		ctx.JSON(500, gin.H{"err": "deu bosta"})
		return
	}
	ctx.JSON(200, gin.H{"data": r})
}
