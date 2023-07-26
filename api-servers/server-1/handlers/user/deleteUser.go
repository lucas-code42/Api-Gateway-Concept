package user

import (
	"net/http"

	"github.com/api-server/lcs42/db"
	"github.com/api-server/lcs42/handlers"
	"github.com/api-server/lcs42/handlers/security"
	"github.com/api-server/lcs42/models"
	"github.com/gin-gonic/gin"
)

func DeleteUser(c *gin.Context) {
	if !security.VerifyToken(c.Request.Header.Get("Authorization")) {
		c.JSON(http.StatusUnauthorized, gin.H{"error": handlers.UNAUTHORIZED})
		return
	}

	user := models.NewUser()
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"error": handlers.BAD_REQUEST})
		return
	}

	if !user.ValidadeUserStruct(models.DELETE_USER) {
		c.JSON(http.StatusBadRequest, gin.H{"error": handlers.BAD_REQUEST})
		return
	}

	rds := db.MountRds()
	defer rds.CloseRds()

	if err := rds.Delete(user); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": handlers.INTERNAL_SERVER_ERROR})
		return
	}

	c.JSON(http.StatusAccepted, gin.H{"userDeletd": user.Id})
}
