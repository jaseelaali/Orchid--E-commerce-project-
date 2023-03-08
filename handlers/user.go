package handlers

import (
	"fmt"
	"github/jaseelaali/orchid/database"
	"github/jaseelaali/orchid/models"
	"github/jaseelaali/orchid/repository"

	//	"net/http"
	//"os"
	"strconv"
	//"time"

	"github.com/gin-gonic/gin"
	//"github.com/golang-jwt/jwt"
	"golang.org/x/crypto/bcrypt"
)

func UserSignUp(r *gin.Context) {

	newUser := models.User{}
	if err := r.Bind(&newUser); err != nil {
		r.JSON(400, gin.H{"message": "error in binding data"})
	}
	Password, err := (bcrypt.GenerateFromPassword([]byte(newUser.Password), 11))
	if err != nil {
		r.JSON(400, gin.H{"message": "error in hashing password"})
	}
	newUser.Password = string(Password)
	newUser.Status = "active"
	err = repository.CreateUser(newUser)
	if err != nil {
		r.JSON(400, err.Error())
		return
	}
	r.JSON(200, gin.H{"success": "Created new user successfully "})
}
func UserLogin(r *gin.Context) {
	var login struct {
		Email    string
		Password string
	}
	if err := r.Bind(&login); err != nil {
		r.JSON(400, gin.H{"message": "error in binding data"})
	}
	//var password string
	user := &models.User{}
	database.DB.Where(&models.User{Email: login.Email, Status: "active"}).First(&user)
	//database.DB.Raw("SELECT password FROM users WHERE email='$1'AND status='active';", login.Email).Scan(&password)
	password := user.Password
	fmt.Println(password)
	//fmt.Println("*****1***")

	err := bcrypt.CompareHashAndPassword([]byte(password), []byte(login.Password))
	// fmt.Println("*********************************** \n ")
	// fmt.Println(password)
	// fmt.Println("*********************************** \n ")
	// fmt.Println(login.Password)

	if err != nil {
		r.JSON(200, gin.H{"message": err.Error()})
		return
	}
	r.JSON(400, gin.H{"message": "login successfully"})
	/*
		//generate jwt token
		token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
			"sub": login.Email,
			"exp": time.Now().Add(time.Hour * 24 * 30).Unix(),
		})

		//sign and get the complete encoded token as a string using the secret key
		tokenstring, err := token.SignedString([]byte(os.Getenv("SECRET")))
		if err != nil {
			r.JSON(400, gin.H{"message": "unable to create token"})
		}
		//send it back
		r.SetSameSite(http.SameSiteLaxMode)
		r.SetCookie("Authorization", tokenstring, 3600*24*30, "", "", false, true)
		r.JSON(200, gin.H{
			//"token": tokenstring
		})

		// func Validate(r *gin.Context){
		// 	r.JSON(200,gin.H{
		// 		"message":"loged in"
		// 	})
		// }
	*/
}

func ViewUser(r *gin.Context) {

	users, err := repository.View()
	if err != nil {
		r.JSON(400, gin.H{
			"message": err.Error()})
	} else {
		r.JSON(200, gin.H{
			//"message":"the list
			"message": users})
	}
}
func BlockUser(r *gin.Context) {
	ID, _ := strconv.Atoi(r.Query("id"))
	err := repository.BlockUser(ID)
	if err != nil {
		r.JSON(400, gin.H{"error": err.Error()})
		return
	}
	r.JSON(200, gin.H{"success": "Blocked user successfully"})
}

func UnBlockUser(r *gin.Context) {
	ID, _ := strconv.Atoi(r.Query("id"))
	err := repository.UnBlockUser(ID)
	if err != nil {
		r.JSON(400, gin.H{"error": err.Error()})
	} else {
		r.JSON(200, gin.H{"success": "Unblocked user successfully"})
	}
}

func BlockedUsers(r *gin.Context) {

	users, err := repository.BlocUsers()
	if err != nil {
		r.JSON(400, gin.H{
			"message": err.Error()})
	} else {
		r.JSON(200, gin.H{
			//"message":"the list
			"message": users})
	}
}
func ActiveUsers(r *gin.Context) {
	users, err := repository.ActiveUser()
	if err != nil {
		r.JSON(400, gin.H{
			"message": err.Error()})
	} else {
		r.JSON(200, gin.H{
			"message": users})
	}
}
func ChangePassword(r *gin.Context) {
	var Password struct {
		//sEmail        string
		Phone_Number string
		NewPassword  string
	}
	if err := r.Bind(&Password); err != nil {
		r.JSON(400, gin.H{"message": "error in binding data"})
	}
	var email string
	//database.DB.Where(&models.User{Email: Password.Email, Phone_Number: Password.Phone_Number}).First(&Id)

	database.DB.Raw("SELECT email FROM user WHERE  phone_number=$1;", Password.Phone_Number).Scan(&email)

	//Id, _ := strconv.Atoi(r.Query("ID"))
	fmt.Println("-------------------------")
	fmt.Println(email)
	fmt.Println("-------------------------")

	err := repository.NewPassword(Password.NewPassword, email)
	if err != nil {
		r.JSON(400, gin.H{
			"message": err.Error})
	} else {
		r.JSON(200, gin.H{
			"message": "password changed"})
	}
}
