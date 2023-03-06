package handlers

import (
	"fmt"
	"github/jaseelaali/orchid/database"
	"github/jaseelaali/orchid/models"
	"github/jaseelaali/orchid/repository"

	"github.com/gin-gonic/gin"
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
	var password string
	database.DB.Raw("SELECT password FROM users WHERE email=$1;", login.Email).Scan(&password)
	fmt.Println(password)
	fmt.Println("login password", login.Password)
	err := bcrypt.CompareHashAndPassword([]byte(password), []byte(login.Password))
	fmt.Println(err)
	if err != nil {
		r.JSON(400, gin.H{
			"message": err.Error(),
		})
	} else {

		r.JSON(200, gin.H{"message": "login successfully"})
	}
}
