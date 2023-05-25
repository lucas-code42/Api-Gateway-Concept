package handlers

import (
	"errors"
	"fmt"
	"net/http"
	"time"

	"github.com/api-server/lcs42/config"
	"github.com/api-server/lcs42/models"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

// DeliveryToken returns a jwt token
func DeliveryToken(c *gin.Context) {
	jwt, err := generateServerToken()
	if err != nil {
		c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}
	c.JSON(http.StatusOK, map[string]string{"token": jwt})
}

// generateServerToken creates the token
func generateServerToken() (string, error) {
	var jwtKey = []byte(config.JWT_KEY)

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"exp": time.Now().Add(time.Second * 20).Unix(), // add time minute as .env variable?
		"iss":    config.JWT_ISSUER,
	})

	tokenSrtring, err := token.SignedString(jwtKey)
	if err != nil {
		errMsg := "could not generate server token"
		fmt.Errorf("%s\nError: %v", errMsg, err)
		return "", errors.New(errMsg)
	}
	return tokenSrtring, nil
}

// VerifyToken validate if the token still valid
func VerifyToken(c *gin.Context) {
	// ref: https://pkg.go.dev/github.com/golang-jwt/jwt/v5#example-Parse-Hmac

	var requestJwt models.JwtModel
	if err := c.ShouldBindJSON(&requestJwt); err != nil {
		c.JSON(403, map[string]bool{"auth": false})
	}

	token, err := jwt.Parse(requestJwt.Token, func(token *jwt.Token) (interface{}, error) {
		fmt.Println(token)
		if token.Method.Alg() != "HS256" {
			return nil, errors.New("invalid")
		}
		return []byte(config.JWT_KEY), nil

	})
	// ... error handling
	if err != nil {
		c.JSON(500, map[string]bool{"auth": false})
		return
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		if claims["iss"] == config.JWT_ISSUER {
			c.JSON(200, map[string]interface{}{"auth": token})
			return
		}
	} else {
		c.JSON(403, map[string]bool{"auth": false})
		return
	}

}

// jwt struct
// {
//     "Raw": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJFeHBpcmVzQXQiOjE2ODQ5ODI4NzEsIklzc3VlciI6ImFwaWdhdGV3YXkkJCJ9.kbQ_O0m6TTnFt090JgCBDE8lTjR2RIGMugjgEhX9m20",
//     "Method": {
//         "Name": "HS256",
//         "Hash": 5
//     },
//     "Header": {
//         "alg": "HS256",
//         "typ": "JWT"
//     },
//     "Claims": {
//         "ExpiresAt": 1684982871,
//         "Issuer": "apigateway$$"
//     },
//     "Signature": "kbQ/O0m6TTnFt090JgCBDE8lTjR2RIGMugjgEhX9m20=",
//     "Valid": true
// }
