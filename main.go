package main

import (
	"github/jaseelaali/orchid/database"
	"github/jaseelaali/orchid/routes"

	"github.com/gin-gonic/gin"
)

func init() {
	database.DatabaseConnection()
}
func main() {
	Route := gin.Default()
	routes.User(Route)
	routes.Admin(Route)

	Route.Run(":9080")
}
