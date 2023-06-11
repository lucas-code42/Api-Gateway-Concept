package user

import (
	"net/http"

	"github.com/api-server/lcs42/db"
	"github.com/api-server/lcs42/handlers"
	"github.com/api-server/lcs42/handlers/security"
	"github.com/api-server/lcs42/models"
	"github.com/gin-gonic/gin"
)

func GetUser(c *gin.Context) {
	if !security.VerifyToken((c.Request.Header.Get("Authorization"))) {
		c.JSON(http.StatusUnauthorized, gin.H{"error": handlers.UNAUTHORIZED})
		return
	}

	userId := c.Query("userId")
	user := models.User{Id: userId}
	if !user.ValidadeUserStruct(models.GET_USER) {
		c.JSON(http.StatusBadRequest, gin.H{"error": handlers.BAD_REQUEST})
		return
	}

	rds := db.MountRds()
	defer rds.CloseRds()

	user, err := rds.Get(userId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": handlers.INTERNAL_SERVER_ERROR})
		return
	}

	c.JSON(http.StatusOK, user)
}
