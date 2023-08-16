package middleware

import (
	"Api-Gateway-lcs42/config"
	"Api-Gateway-lcs42/routers/httpHandler"
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
)

/*
	TODO: ADD UM CONTEXTO PARA PEGAR O TEMPO DE PROCESSAMENTO
*/

func DummyMiddleware() gin.HandlerFunc {
	if config.APIGATEWAY_KEY == "" {
		log.Fatal("cannot read environment variable")
	}
	return func(c *gin.Context) {
		authKey := c.Request.Header.Get("Authorization")
		if authKey != config.APIGATEWAY_KEY {
			respondWithError(c, 500, "deu bosta")
			return
		}

		s := c.Param("server")
		if s == "" || (s != "server1" && s != "server2") {
			respondWithError(c, 500, fmt.Sprintf("não enviou o serverName ou enviou errado %s", s))
			return
		}

		jwt, err := httpHandler.GetJwt(s)
		if err != nil {
			respondWithError(c, 500, "mid nao conseguiu pegar o jwt do server")
			return
		}
		c.Set("jwt", jwt.Token)

		c.Next()
	}
}

func respondWithError(c *gin.Context, code int, message interface{}) {
	c.AbortWithStatusJSON(code, gin.H{"error": message})
}
