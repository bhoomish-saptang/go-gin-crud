package crud

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-gin/constants"
	"github.com/go-gin/database"
)

func PostUserDetails(c *gin.Context) {
	var newUser constants.User
	if err := c.ShouldBindJSON(&newUser); err != nil {
		log.Println("error occured in create user details endpoint at BindJSON of constants.User")
		c.IndentedJSON(http.StatusNotAcceptable, gin.H{"message": "Required all fields neccesary and check the JSON Request Body"})
		return
	} else {
		id := newUser.ID
		_, err := database.FindUserDetailsByID(id)
		if err == nil {
			c.IndentedJSON(http.StatusAlreadyReported, gin.H{"message": "the User details ID is already exist"})
		} else {
			err := database.InsertUserData(newUser)
			if err != nil {
				c.IndentedJSON(http.StatusBadRequest, newUser)
			} else {
				c.IndentedJSON(http.StatusCreated, newUser)
			}
		}
	}
}
