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
	r.GET("/test2", test2)
	r.Run(":8080") // listen and serve on 0.0.0.0:8080
}

func test1(c *gin.Context) {
	c.JSON(200, map[string]string{"server1_says": callApi()})
}

func test2(c *gin.Context) {
	c.JSON(200, map[string]string{"server2_says": callApi()})
	c.Header("tls", "9999")

}


func callApi() string {
	url := "http://127.0.0.1:8081/api"
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

	return string(body)
}
