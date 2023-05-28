package routers

import (
	"fmt"
	"io"
	"net/http"

	"github.com/gin-gonic/gin"
)

func sayHi(rg *gin.RouterGroup) {
	healthCheck := rg.Group("/health")

	healthCheck.GET("/", func(ctx *gin.Context) {
		url := "http://127.0.0.1:2001/server1/health"
		method := "GET"

		client := &http.Client{}
		req, err := http.NewRequest(method, url, nil)

		if err != nil {
			fmt.Println(err)
			return
		}
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

		ctx.JSON(200, string(body))
	})

}
