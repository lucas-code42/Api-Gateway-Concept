package middleware

import (
	"Api-Gateway-lcs42/config"
	"log"

	"github.com/gin-gonic/gin"
)

func DummyMiddleware() gin.HandlerFunc {
	if config.APIGATEWAY_JWT_KEY == "" {
		log.Fatal("cannot read environment variable")
	}
	return func(c *gin.Context) {
		authKey := c.Request.Header.Get("Authorization")
		if authKey != config.APIGATEWAY_JWT_KEY {
			respondWithError(c, 500, "deu bosta")
			return
		}
		c.Next()
	}
}

func respondWithError(c *gin.Context, code int, message interface{}) {
	c.AbortWithStatusJSON(code, gin.H{"error": message})
}
