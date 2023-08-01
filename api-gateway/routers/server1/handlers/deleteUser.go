package handlers

import (
	"Api-Gateway-lcs42/config"
	"Api-Gateway-lcs42/routers/server1/tools"
	"bytes"
	"encoding/json"
	"fmt"

	"github.com/gin-gonic/gin"
)

func Deleteuser(ctx *gin.Context) {
	var clientJson interface{}
	if err := ctx.BindJSON(&clientJson); err != nil {
		ctx.JSON(500, gin.H{"err": "deu bosta no delete user"})
		return
	}

	clientData := new(bytes.Buffer)
	json.NewEncoder(clientData).Encode(clientJson)

	fmt.Println(clientData)

	r, err := tools.RequestServer1("DELETE", config.SERVER1_PATH, clientData)
	if err != nil {
		ctx.JSON(500, gin.H{"err": "deu bosta no delete user request"})
		return
	}

	ctx.JSON(200, gin.H{"data": r})
}
