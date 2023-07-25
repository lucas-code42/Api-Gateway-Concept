package handlersServer1

import (
	"Api-Gateway-lcs42/config"
	"Api-Gateway-lcs42/models"
	"Api-Gateway-lcs42/utils"
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func CreateUser(ctx *gin.Context) {
	start := time.Now()

	url := fmt.Sprintf("%s/%s", config.DEFAULT_HOST_SERVER1, config.SERVER1_CREATE_PATH)
	method := "POST"

	var clientJson interface{}
	if err := ctx.BindJSON(&clientJson); err != nil {
		ctx.JSON(500, gin.H{"err": "deu bosta"})
		return
	}
	clientData := new(bytes.Buffer)
	json.NewEncoder(clientData).Encode(clientJson)

	client := &http.Client{}
	req, err := http.NewRequest(method, url, clientData)
	if err != nil {
		ctx.JSON(500, gin.H{"err": "error"})
	}

	jwt, err := GetJwt("server1")
	fmt.Println(jwt)
	if err != nil {
		ctx.JSON(500, gin.H{"err": "error"})
	}
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Authorization", jwt.Token)

	res, err := client.Do(req)
	if err != nil {
		ctx.JSON(500, gin.H{"err": "error"})
		return
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		ctx.JSON(500, gin.H{"err": "could not read server1 response"})
		return
	}

	data, err := utils.ParseDtoResponse(body)
	if err != nil {
		ctx.JSON(500, gin.H{"err": "could not mount dto response"})
		return
	}

	var buff []interface{}
	buff = append(buff, data)
	response := models.DtoResponse{
		Message:       "msg",
		Id:            uuid.NewString(),
		Data:          buff,
		StatusCode:    200,
		ExecutionTime: time.Duration(time.Since(start).Milliseconds()),
	}

	ctx.JSON(200, gin.H{"data": response})
}
