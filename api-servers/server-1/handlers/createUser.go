package handlers

import (
	"log"
	"net/http"

	"github.com/api-server/lcs42/db"
	"github.com/api-server/lcs42/models"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

// CreateUser is a gin handler that will call all methods responsable for create user
func CreateUser(c *gin.Context) {
	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusUnprocessableEntity, map[string]string{"err": "could not create user"})
		return
	}
	user.Id = uuid.NewString()

	rds := db.MountRds()
	defer rds.CloseRds()

	if err := rds.Create(user); err != nil {
		c.JSON(http.StatusInternalServerError, map[string]string{"err": "could not create user"})
		log.Println("Error >>>>>>>>>", err)
		// log.Fatal(err)
		return
	}

	c.JSON(http.StatusCreated, user)
}
