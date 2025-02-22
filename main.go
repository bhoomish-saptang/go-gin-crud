package main

import (
	"os"

	"github.com/gin-gonic/gin"
	"github.com/go-gin/authentication"
	"github.com/go-gin/config"
	"github.com/go-gin/crud"
	"github.com/go-gin/database"
	"github.com/go-gin/logger"
	"github.com/go-gin/middleware"
)

func init() {
	logger.InitLogger()
	config.SetEnvironmentVariables()
	database.ConnectMongodb()
}
func main() {
	logger.Log.Println("Your server is ready")

	route := gin.Default()
	//userDetails CRUD API endpoints
	route.POST("/createUserDetails", crud.PostUserDetails)
	route.GET("/getAllUserDetails", crud.GetAllUserDetails)
	route.GET("/getUserDetailsById/:id", crud.GetUserDetailsById)
	route.PUT("/updateUserDetailsById/:id", crud.UpdateUserDetailsById)
	route.DELETE("/deleteUserDetailsById/:id", crud.DeleteUserDetailsByID)

	//Authenticaton Endpoints
	route.POST("/createAuthUser", authentication.CreateUser)
	route.POST("/authUserLogin", authentication.AuthUserLogin)
	route.GET("/getAuthUser", middleware.CheckAuth, crud.GetAuthUser)
	route.Run(os.Getenv("PORT"))

}
