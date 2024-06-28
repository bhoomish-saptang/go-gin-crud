package crud

import "github.com/gin-gonic/gin"

func GetAuthUser(c *gin.Context) {

	user, _ := c.Get("currentUser")

	c.JSON(200, gin.H{
		"user": user,
	})
}
