package user

import (
	"log"
	"net/http"

	"github.com/api-server/lcs42/db"
	"github.com/api-server/lcs42/handlers/security"
	"github.com/api-server/lcs42/models"
	"github.com/gin-gonic/gin"
)

func UpdateUser(c *gin.Context) {
	if !security.VerifyToken((c.Request.Header.Get("Authorization"))) {
		c.JSON(http.StatusUnauthorized, gin.H{"err": "unauthorized"})
		return
	}

	var updatedUser models.User
	if err := c.ShouldBindJSON(&updatedUser); err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"err": "could not update user"})
		return
	}

	rds := db.MountRds()
	defer rds.CloseRds()

	if err := rds.Update(updatedUser); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"err": "could not update user"})
		log.Println("Error >>>>>>>>>", err)
		// log.Fatal(err)
		return
	}

	c.JSON(http.StatusCreated, updatedUser)

}
