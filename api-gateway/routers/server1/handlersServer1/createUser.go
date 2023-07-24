package handlersServer1

import (
	"Api-Gateway-lcs42/config"
	"fmt"
	"io"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func CreateUser(ctx *gin.Context) {
	url := fmt.Sprintf("%s/%s", config.DEFAULT_HOST_SERVER1, config.SERVER1_CREATE_PATH)
	method := "POST"

	payload := strings.NewReader(`{
		"name": "linux",
		"email": "linux@email.com"
	}`)

	client := &http.Client{}
	req, err := http.NewRequest(method, url, payload)
	if err != nil {
		ctx.JSON(503, gin.H{"err": "error"})
	}

	jwt, err := GetJwt("server1")
	fmt.Println(jwt)
	if err != nil {
		ctx.JSON(503, gin.H{"err": "error"})
	}
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Authorization", jwt.Token)

	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(body))
}
