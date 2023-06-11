package handlers

import (
	"fmt"
	"net/http"

	"github.com/api-server/lcs42/config"
	"github.com/api-server/lcs42/db"
	"github.com/gin-gonic/gin"
)

func GetUser(c *gin.Context) {
	if c.Request.Header.Get("Authorization") != config.SECURITY_KEY {
		c.JSON(http.StatusInternalServerError, gin.H{"err": "unauthorized"})
		return
	}

	if !verifyToken(c.Request.Header.Get("Authorization")) {
		c.JSON(http.StatusUnauthorized, gin.H{"err": "unauthorized"})
		return
	}

	userId := c.Query("userId")

	rds := db.MountRds()
	defer rds.CloseRds()

	user, err := rds.Get(userId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "cannot get user"})
		return
	}

	fmt.Println(user)

	c.JSON(http.StatusOK, user)
}
