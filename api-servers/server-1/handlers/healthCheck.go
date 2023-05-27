package handlers

import (
	"time"

	"github.com/gin-gonic/gin"
)

// HealthCheck just health checker
func HealthCheck(c *gin.Context) {
	c.JSON(200, gin.H{"server1": time.Now().Format("2017.09.07 17:06:06")})
}
