package handlers

import (
	"github/jaseelaali/orchid/database"
	"github/jaseelaali/orchid/models"
	"github/jaseelaali/orchid/repository"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

func AdminLogin(r *gin.Context) {
	var login struct {
		Email    string
		Password string
	}
	if err := r.Bind(&login); err != nil {
		r.JSON(400, gin.H{"message": "error in binding data"})
	}
	//var password string
	admin := models.Admin{}
	database.DB.Where("email=?", login.Email).First(&admin)
	//err := database.DB.Raw("SELECT password from admins where email='$1'", login.Email).Scan(&password).Error
	err := bcrypt.CompareHashAndPassword([]byte(admin.Password), []byte(login.Password))
	if err != nil {
		r.JSON(400, gin.H{
			"message": "passwords or email are not matching",
		})
		return
	}
	/*-----------------------------------------------------------------------------------------------------------------------------*/
	/*password, err := bcrypt.GenerateFromPassword(([]byte(login.Password)), 11)
	Password := string(password)
	fmt.Printf("****password want to databse:%v", Password)
	database.DB.Raw("UPDATE admins SET password=$1 WHERE email=$2;", Password, admin.Email).Scan(&models.Admin{})*/

	/*---------------------------------------------------------------------------------------------------------------------------------*/
	//generate jwt token
	//here call token function
	token, err := repository.Token(login.Email)

	//sign and get the complete encoded token as a string using the secret key
	tokenstring, err := token.SignedString([]byte(os.Getenv("SECRET")))
	if err != nil {
		r.JSON(400, gin.H{"message": "unable to create token"})
		return
	}
	// //send it back
	r.SetSameSite(http.SameSiteLaxMode)
	r.SetCookie("Authorization", tokenstring, 3600*24*30, "", "", false, true)
	r.JSON(200, gin.H{
		"token":   tokenstring,
		"message": "login successfully",
	})

	// func Validate(r *gin.Context){
	// 	r.JSON(200,gin.H{
	// 		"message":"loged in"
	// 	})
	// }

}
