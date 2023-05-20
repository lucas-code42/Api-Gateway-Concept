package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.New()
	r.GET("/server1", test1)
	r.Run(":2001") // listen and serve on 0.0.0.0:8080
}

func test1(c *gin.Context) {
	c.Header("token", "server1")
	c.JSON(200, gin.H{"server1": "hello i'm server 1 :)"})
}
