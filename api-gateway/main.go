package main

import (
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	
)

func main() {
	r := gin.New()
	r.GET("/test1", test1)
	r.Run(":2004") // listen and serve on 0.0.0.0:8080
}

func test1(c *gin.Context) {
	response, header := callApi()
	c.Header("ApiGateWayToken", "API-GATEWAY-TOKEN")
	c.Header("Server1Token", header)
	c.JSON(200, map[string]string{"server1_says": response})
}

func callApi() (string, string) {
	url := "http://127.0.0.1:2001/server1"
	method := "GET"

	client := &http.Client{}
	req, err := http.NewRequest(method, url, nil)

	if err != nil {
		fmt.Println(err)
		log.Fatal(err)
	}
	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		log.Fatal(err)
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		log.Fatal(err)
	}

	fmt.Println()
	return string(body), res.Header["Token"][0]
}
