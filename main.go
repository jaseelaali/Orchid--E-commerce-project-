package main

import (
	"github/jaseelaali/orchid/controllers"
	"github/jaseelaali/orchid/database"

	"github.com/gin-gonic/gin"
)

func init() {
	database.DatabaseConnection()
}
func main() {
	routes := gin.Default()
	routes.GET("/admin", controllers.Admin)
	routes.GET("/user", controllers.User)
	routes.Run("localhost:8080")
}
