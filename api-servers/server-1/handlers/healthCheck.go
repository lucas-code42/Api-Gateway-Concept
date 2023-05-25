package handlers

import (
	"time"

	"github.com/gin-gonic/gin"
)

func HealthCheck(c *gin.Context) {
	c.JSON(200, map[string]string{"server1": time.Now().Format("2017.09.07 17:06:06")})
}
