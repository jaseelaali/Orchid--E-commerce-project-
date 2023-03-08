package main

import (
	"github/jaseelaali/orchid/database"
	"github/jaseelaali/orchid/handlers"

	"github.com/gin-gonic/gin"
)

func init() {
	database.DatabaseConnection()
}
func main() {
	routes := gin.Default()
	routes.GET("/admin", handlers.Admin)
	routes.POST("/usersignup", handlers.UserSignUp)
	routes.POST("/userlogin", handlers.UserLogin)
	//routes.POST("/re", middleware .RequiredAuthentication, handlers.Validate)
	routes.GET("/viewuser", handlers.ViewUser)
	routes.POST("/blockuser", handlers.BlockUser)
	routes.POST("/unblockuser", handlers.UnBlockUser)
	routes.GET("/viewblockedusers", handlers.BlockedUsers)
	routes.GET("/viewactiveusers", handlers.ActiveUsers)
	routes.POST("/changepassword", handlers.ChangePassword)
	routes.Run(":8080")
}
