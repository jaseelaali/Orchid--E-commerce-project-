package routes

import (
	"github/jaseelaali/orchid/handlers"
	"github/jaseelaali/orchid/middleware"

	//"github/jaseelaali/orchid/middleware"

	"github.com/gin-gonic/gin"
)

func User(r *gin.Engine) {
	r.LoadHTMLGlob("template/*.html")
	user := r.Group("/user")

	{
		// user signup+login
		user.POST("signup", handlers.UserSignUp)
		user.POST("login", handlers.UserLogin)
		//profile
		user.GET("/aboutme", middleware.RequiredAuthenticationUser, handlers.Profile)
		user.PATCH("/editmyprofile", middleware.RequiredAuthenticationUser, handlers.EditProfile)
		user.DELETE("/deleteprofile",middleware.RequiredAuthenticationUser,handlers.DeleteProfile)
		// cart management
		user.POST("addcart", middleware.RequiredAuthenticationUser, handlers.AddCart)
		user.GET("viewcart", middleware.RequiredAuthenticationUser, handlers.ViewCart)
		user.DELETE("deleteitem", middleware.RequiredAuthenticationUser, handlers.DeleteItem)
		//change password
		user.POST("changepassword", middleware.RequiredAuthenticationUser, handlers.ChangePassword)
		user.POST("verifyotp", middleware.RequiredAuthenticationUser, handlers.VerifyOtp)
		// address
		user.POST("/addaddress", middleware.RequiredAuthenticationUser, handlers.Address)
		user.PATCH("/editaddress", middleware.RequiredAuthenticationUser, handlers.EditAddress)
		user.DELETE("/deleteaddress", middleware.RequiredAuthenticationUser, handlers.DeleteAddress)
		user.GET("/viewaddress", middleware.RequiredAuthenticationUser, handlers.ViewAddress)
		// order management

		//payment
		user.GET("/razorpay", handlers.Razorpay)
		user.GET("/payment-success", handlers.Payment_Success)
	}
}
func Admin(r *gin.Engine) {
	admin := r.Group("/admin")
	{
		//admin login
		admin.POST("login", handlers.AdminLogin)
		//user management
		admin.GET("view", middleware.RequiredAuthenticationAdmin, handlers.ViewUser)
		admin.POST("block", middleware.RequiredAuthenticationAdmin, handlers.BlockUser)
		admin.POST("unblock", middleware.RequiredAuthenticationAdmin, handlers.UnBlockUser)
		admin.GET("viewblockedusers", middleware.RequiredAuthenticationAdmin, handlers.BlockedUsers)
		admin.GET("viewunblockedusers", middleware.RequiredAuthenticationAdmin, handlers.ActiveUsers)
		// category mangement
		admin.POST("addcategory", middleware.RequiredAuthenticationAdmin, handlers.AddCategory)
		admin.PATCH("editcategory", middleware.RequiredAuthenticationAdmin, handlers.EditCategory)
		admin.DELETE("deletecategory", middleware.RequiredAuthenticationAdmin, handlers.DeleteCategory)
		admin.GET("viewcategory", middleware.RequiredAuthenticationAdmin, handlers.ViewCategory)
		// sub category management
		admin.POST("addsubcategory", middleware.RequiredAuthenticationAdmin, handlers.AddSubCategory)
		admin.PATCH("editsubcategory", middleware.RequiredAuthenticationAdmin, handlers.EditSubCategory)
		admin.DELETE("deletesubcategory", middleware.RequiredAuthenticationAdmin, handlers.DeleteSubCategory)
		admin.GET("viewsubcategory", handlers.ViewSubCategory)
		//product managemant
		admin.POST("addproduct", middleware.RequiredAuthenticationAdmin, handlers.AddProducts)
		admin.PATCH("editproducts", middleware.RequiredAuthenticationAdmin, handlers.EditProducts)
		admin.DELETE("deleteproducts", middleware.RequiredAuthenticationAdmin, handlers.DeleteProducts)
		admin.GET("viewproducts", middleware.RequiredAuthenticationAdmin, handlers.ViewProducts)
		// cart management
	}
}
