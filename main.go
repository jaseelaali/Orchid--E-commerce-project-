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
	routes.Run(":8080")
}
