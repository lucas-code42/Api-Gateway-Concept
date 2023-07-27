package handlers

import (
	"Api-Gateway-lcs42/config"
	"Api-Gateway-lcs42/utils"
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

// TODO: improve the default response
// think of a way how to use dto model
// this request is to long...
func HealthCheck(ctx *gin.Context) {
	url := fmt.Sprintf("%s/health", config.DEFAULT_HOST_SERVER1)
	method := "GET"

	client := &http.Client{}
	req, err := http.NewRequest(method, url, nil)

	if err != nil {
		log.Println(err)
		ctx.JSON(http.StatusInternalServerError, gin.H{"err": "error"})
		return
	}

	res, err := client.Do(req)
	if err != nil {
		log.Println(err)
		ctx.JSON(http.StatusInternalServerError, gin.H{"err": "error"})
		return
	}
	defer res.Body.Close()
	if res.StatusCode != 200 {
		ctx.JSON(
			http.StatusServiceUnavailable,
			gin.H{"err": "the requested service is currently unavailable. Please try again later."},
		)
		return
	}

	body, err := io.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return
	}

	var response interface{}
	response, err = utils.ParseDtoResponse(body)
	if err != nil {
		ctx.JSON(
			http.StatusServiceUnavailable,
			gin.H{"err": "the requested service is currently unavailable. Please try again later."},
		)
		return
	}

	ctx.JSON(200, gin.H{"data": response})
}
