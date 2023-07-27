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

	userId := c.Query("id")
	_, err := uuid.Parse(userId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": handlers.BAD_REQUEST})
		return
	}

	rds := db.MountRds()
	defer rds.CloseRds()
	if rds.Exist(userId) != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": handlers.BAD_REQUEST})
		return
	}

	if err := rds.Delete(userId); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": handlers.INTERNAL_SERVER_ERROR})
		return
	}

	c.JSON(http.StatusAccepted, gin.H{"userDeletd": userId})
}
