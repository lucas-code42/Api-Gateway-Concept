package user

import (
	"net/http"

	"github.com/api-server/lcs42/db"
	"github.com/api-server/lcs42/handlers"
	"github.com/api-server/lcs42/handlers/security"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func DeleteUser(c *gin.Context) {
	if !security.VerifyToken(c.Request.Header.Get("Authorization")) {
		c.JSON(http.StatusUnauthorized, gin.H{"error": handlers.UNAUTHORIZED})
		return
	}

	type user_id struct {
		Id string `json:"id"`
	}
	var user user_id
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": handlers.BAD_REQUEST})
		return
	}

	_, err := uuid.Parse(user.Id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": handlers.BAD_REQUEST})
		return
	}

	rds := db.MountRds()
	defer rds.CloseRds()
	if rds.Exist(user.Id) != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": handlers.BAD_REQUEST})
		return
	}

	if err := rds.Delete(user.Id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": handlers.INTERNAL_SERVER_ERROR})
		return
	}

	c.JSON(http.StatusAccepted, gin.H{"userDeletd": user})
}
