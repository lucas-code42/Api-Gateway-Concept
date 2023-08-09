package handlers

import (
	"Api-Gateway-lcs42/routers/tools"
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func ServerInterfacePost(c *gin.Context) {
	start := time.Now()

	// ? PrepareRequest deveria lidar com o paylod?
	var clientJson interface{}
	if err := c.BindJSON(&clientJson); err != nil {
		c.JSON(400, gin.H{"err": "unprocessable entity"})
		return
	}
	clientData := new(bytes.Buffer)
	json.NewEncoder(clientData).Encode(clientJson)

	serverHost, err := PrepareRequest(c, POST, c.Keys["jwt"])
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "internal error"})
		return
	}

	r, err := tools.DoRequest(
		serverHost.Url, serverHost.Path, fmt.Sprintf("%v", c.Keys["jwt"]), "POST", clientData,
	)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "internal error"})
		return
	}

	r.ExecutionTime = time.Duration(time.Since(start).Milliseconds())
	c.JSON(http.StatusOK, gin.H{"data": r})
}
