package routes

import (
	"github/jaseelaali/orchid/handlers"
	"github/jaseelaali/orchid/middleware"

	"github.com/gin-gonic/gin"
)

func User(r *gin.Engine) {
	user := r.Group("/user")
	{
		// user signup+login
		user.POST("signup", handlers.UserSignUp)
		user.POST("login", handlers.UserLogin)
		// cart management
		user.POST("addcart", middleware.RequiredAuthentication, handlers.AddCart)
		user.GET("viewcart", middleware.RequiredAuthentication, handlers.ViewCart)
		user.DELETE("deleteitem", middleware.RequiredAuthentication, handlers.DeleteItem)
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
		// category mangement
		admin.POST("addcategory", handlers.AddCategory)
		// sub category management
		admin.POST("addsubcategory", handlers.AddSubCategory)
		//product managemant
		admin.POST("addproduct", handlers.AddProducts)
		// cart management
	}
}
