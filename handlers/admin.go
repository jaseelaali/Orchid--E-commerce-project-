package handlers

import (
	"fmt"
	"github/jaseelaali/orchid/database"
	"github/jaseelaali/orchid/models"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
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
	err := database.DB.Where("email=?", login.Email).First(&admin).Error
	//err := database.DB.Raw("SELECT password from admins where email='$1'", login.Email).Scan(&password).Error
	password := admin.Password
	fmt.Println(password)
	fmt.Println("****1****")
	fmt.Println(login.Password)
	fmt.Println("****2****")
	fmt.Println(password)
	fmt.Println("****3****")

	//err := bcrypt.CompareHashAndPassword([]byte(password), []byte(login.Password))
	if login.Password == password {
		r.JSON(200, gin.H{"message": "login successfully"})

	}
	r.JSON(400, gin.H{"message": err.Error()})
	return

	//generate jwt token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": login.Email,
		"exp": time.Now().Add(time.Hour * 24 * 30).Unix(),
	})

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
		"token": tokenstring,
	})

	// func Validate(r *gin.Context){
	// 	r.JSON(200,gin.H{
	// 		"message":"loged in"
	// 	})
	// }

}
