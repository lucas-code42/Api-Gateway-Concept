package routers

import (
	"Api-Gateway-lcs42/config"
	"Api-Gateway-lcs42/routers/server1"
	"Api-Gateway-lcs42/routers/server2"
	"Api-Gateway-lcs42/routers/tools"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

var httpEngine = gin.Default()

// Run initialize api
func Run() {
	mountRoutes()
	err := httpEngine.Run(fmt.Sprintf(":%s", config.PORT))
	if err != nil {
		log.Fatal("Could not start APIGATEWAY")
	}
}

func mountRoutes() {
	/*
		! routers will be replaced
	*/
	SERVER_1_ROUTE_GROUP := httpEngine.Group(fmt.Sprintf("%s/server1", config.SERVER_DEFAULT_PATH))
	server1.Server1(SERVER_1_ROUTE_GROUP)

	SERVER_2_ROUTE_GROUP := httpEngine.Group(fmt.Sprintf("%s/server2", config.SERVER_DEFAULT_PATH))
	server2.Server2(SERVER_2_ROUTE_GROUP)

	// TODO: modular uma interface padrão que atenda todos os servers...
	httpEngine.GET(fmt.Sprintf("%s/:serverName/", config.SERVER_DEFAULT_PATH), serverGetInterface)
}

func serverGetInterface(c *gin.Context) {
	start := time.Now()

	server := c.Param("serverName")

	var url string
	var path string

	switch server {
	case "server1":
		url = config.DEFAULT_HOST_SERVER1
		path = config.DEFAULT_HOST_SERVER1
	case "server2":
		url = config.DEFAULT_HOST_SERVER2
		path = config.DEFAULT_HOST_SERVER2
	default:
		c.JSON(http.StatusNotFound, gin.H{"error": "unknown server"})
	}

	jwtToken = tools.GetJwt()

	r, err := tools.GetRequest(url, path)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "internal error"})
	}
	r.ExecutionTime = time.Duration(time.Since(start).Milliseconds())

	c.JSON(http.StatusOK, gin.H{"data": r})
}
