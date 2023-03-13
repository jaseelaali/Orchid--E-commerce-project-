package routes

import (
	"github/jaseelaali/orchid/handlers"

	"github.com/gin-gonic/gin"
)

func User(r *gin.Engine) {
	user := r.Group("/user")
	{
		// user signup+login
		user.POST("signup", handlers.UserSignUp)
		user.POST("login", handlers.UserLogin)
	}
}
func Admin(r *gin.Engine) {
	admin := r.Group("/admin")
	{
		//admin login
		admin.POST("login", handlers.AdminLogin)
		//user management
		admin.GET("view", handlers.ViewUser)
		admin.POST("block", handlers.BlockUser)
		admin.POST("unblock", handlers.UnBlockUser)
		admin.GET("viewblockedusers", handlers.BlockedUsers)
		admin.GET("viewunblockedusers", handlers.ActiveUsers)
		//product managemant
		admin.POST("add", handlers.AddProducts)

	}
}
