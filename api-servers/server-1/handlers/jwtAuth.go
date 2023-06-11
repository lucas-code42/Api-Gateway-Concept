package handlers

import (
	"errors"
	"fmt"
	"net/http"
	"time"

	"github.com/api-server/lcs42/config"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

var JWT_EXP = 30

// DeliveryToken returns a jwt token
func DeliveryToken(c *gin.Context) {
	if c.Request.Header.Get("Authorization") != config.SECURITY_KEY {
		c.JSON(http.StatusInternalServerError, gin.H{"err": "invalid permission"})
		return
	}

	jwt, err := generateServerToken()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}
	c.JSON(http.StatusOK, gin.H{"token": jwt})
}

// generateServerToken creates the token
func generateServerToken() (string, error) {
	var jwtKey = []byte(config.JWT_KEY)

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"exp": time.Now().Add(time.Minute * time.Duration(JWT_EXP)).Unix(),
		"iss": config.JWT_ISSUER,
	})

	tokenSrtring, err := token.SignedString(jwtKey)
	if err != nil {
		errMsg := "could not generate server token"
		fmt.Printf("%s\nError: %v", errMsg, err)
		return "", errors.New(errMsg)
	}
	return tokenSrtring, nil
}

// VerifyToken validate if the token still valid
func verifyToken(JwtToken string) bool {
	// ref: https://pkg.go.dev/github.com/golang-jwt/jwt/v5#example-Parse-Hmac

	token, err := jwt.Parse(JwtToken, func(token *jwt.Token) (interface{}, error) {
		if token.Method.Alg() != "HS256" {
			return nil, errors.New("invalid")
		}
		return []byte(config.JWT_KEY), nil

	})
	// ... error handling
	if err != nil {
		return false
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		if claims["iss"] == config.JWT_ISSUER {
			return true
		}
	} else {
		return false
	}

	return false
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
