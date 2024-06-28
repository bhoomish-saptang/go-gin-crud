package authentication

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-gin/constants"
	"github.com/go-gin/database"
	"go.mongodb.org/mongo-driver/bson"
	"golang.org/x/crypto/bcrypt"
)

func AuthUserLogin(c *gin.Context) {

	var authInput constants.AuthInput

	if err := c.ShouldBindJSON(&authInput); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var userFound constants.AuthUser
	filter := bson.M{"username": authInput.Username}
	database.CollectionAuth.FindOne(context.Background(), filter).Decode(&userFound)
	if userFound.Username == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "username does not exit,create one"})
		return
	}

	if err := bcrypt.CompareHashAndPassword([]byte(userFound.Password), []byte(authInput.Password)); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid password"})
		return
	}

	c.IndentedJSON(http.StatusAccepted, userFound)

}
