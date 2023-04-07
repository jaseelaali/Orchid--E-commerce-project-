package handlers

import (
	"fmt"
	"github/jaseelaali/orchid/database"
	"github/jaseelaali/orchid/models"
	"github/jaseelaali/orchid/repository"
	"net/http"
	"os"

	"strconv"

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
		return
	}
	//var password string
	user := &models.User{}
	database.DB.Where(&models.User{Email: login.Email, Status: "active"}).First(&user)
	fmt.Println()
	if user == nil {
		r.JSON(400, gin.H{
			"message": "innalid user",
		})
		return
	}

	//database.DB.Raw("SELECT password FROM users WHERE email='$1'AND status='active';", login.Email).Scan(&password)
	password := user.Password

	err := bcrypt.CompareHashAndPassword([]byte(password), []byte(login.Password))

	if err != nil {
		r.JSON(400, gin.H{"message": "passwords are not matching "})
		return
	}
	//r.JSON(200, gin.H{"message": "login successfully"})

	//generate jwt token
	//here call token function
	token, err := repository.Token(login.Email)

	//sign and get the complete encoded token as a string using the secret key
	tokenstring, err := token.SignedString([]byte(os.Getenv("SECRET")))
	if err != nil {
		r.JSON(400, gin.H{"message": "unable to create token"})
		return
	}
	//send it back
	r.SetSameSite(http.SameSiteLaxMode)
	r.SetCookie("Authorization", tokenstring, 3600*24*30, "", "", false, true)
	r.JSON(200, gin.H{
		"token":   tokenstring,
		"message": "login successfully",
	})

}

func ViewUser(r *gin.Context) {

	page:=r.Query("page")
	if page==""{
		r.JSON(400,gin.H{
			"message":"didn't get page number",
		})
		return
	}
	perpage:=r.Query("perpage")
	if perpage==""{
		r.JSON(400,gin.H{
			"message":"didn't get perpage number",
		})
		return
	}
	pagenumber,_:=strconv.Atoi(page)
	perpagenumber,_:=strconv.Atoi(perpage)

	users, metaData, err := repository.View(pagenumber,perpagenumber)
	if err != nil {
		r.JSON(400, gin.H{
			"message": err.Error()})
	} else {
		r.JSON(200, gin.H{
			//"message":"the list
			"List of users": users,
			"metadata":      metaData})
	}
}
func SpeacificUser(r *gin.Context) {
	var body struct {
		Id           int    `json:"id"`
		User_Name    string `json:"user_name"`
		Email        string `json:"email"`
		Phone_Number string `json:"phone_number"`
	}
	if err := r.Bind(&body); err != nil {
		r.JSON(400, gin.H{"message": "error in binding data"})
		return
	}
	fmt.Printf(".................%v", body.Id)
	if body.Id != 0 {
		information, err := repository.UserById(body.Id)

		if err != nil {
			r.JSON(400, gin.H{
				"message": err.Error(),
			})
			return
		}
		r.JSON(200, gin.H{
			"success": information,
		})
	}
	if body.User_Name != "" {
		information, err := repository.UserByName(body.User_Name)
		if err != nil {
			r.JSON(400, gin.H{
				"message": err.Error(),
			})
			return
		}
		r.JSON(200, gin.H{
			"success": information,
		})
	}
	if body.Email != "" {
		information, err := repository.UserByEmail(body.Email)
		if err != nil {
			r.JSON(400, gin.H{
				"message": err.Error(),
			})
			return
		}
		r.JSON(200, gin.H{
			"success": information,
		})
	}
	if body.Phone_Number != "" {
		information, err := repository.UserByNumber(body.Phone_Number)
		if err != nil {
			r.JSON(400, gin.H{
				"message": err.Error(),
			})
			return
		}
		r.JSON(200, gin.H{
			"success": information,
		})
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
	var Body struct {
		Page    int `json:"page" binding:"required"`
		Perpage int `json:"perpage" binding:"required"`
	}
	err := r.ShouldBind(&Body)
	if err != nil {
		r.JSON(400, gin.H{
			"message": err.Error()})
		return
	}

	users, metadata, err := repository.BlocUsers(Body.Page, Body.Perpage)
	if err != nil {
		r.JSON(400, gin.H{
			"message": err.Error()})
	} else {
		r.JSON(200, gin.H{
			//"message":"the list
			"message":  users,
			"metadata": metadata})
	}

}
func ActiveUsers(r *gin.Context) {
	var Body struct {
		Page    int `json:"page" binding:"required"`
		Perpage int `json:"perpage" binding:"required"`
	}
	err := r.ShouldBind(&Body)
	if err != nil {
		r.JSON(400, gin.H{
			"message": err.Error()})
		return
	}

	users, metadata, err := repository.ActiveUser(Body.Page, Body.Perpage)
	if err != nil {
		r.JSON(400, gin.H{
			"message": err.Error()})
	} else {
		r.JSON(200, gin.H{
			"message":  users,
			"metadata": metadata})
	}
}

// func ChangePassword(r *gin.Context) {
// 	var Password struct {
// 		//sEmail        string
// 		Phone_Number string
// 		NewPassword  string
// 	}
// 	if err := r.Bind(&Password); err != nil {
// 		r.JSON(400, gin.H{"message": "error in binding data"})
// 	}
// 	var email string
// 	//database.DB.Where(&models.User{Email: Password.Email, Phone_Number: Password.Phone_Number}).First(&Id)

// 	database.DB.Raw("SELECT email FROM user WHERE  phone_number=$1;", Password.Phone_Number).Scan(&email)

// 	//Id, _ := strconv.Atoi(r.Query("ID"))
// 	fmt.Println("-------------------------")
// 	fmt.Println(email)

// 	fmt.Println("-------------------------")

// 	err := repository.NewPassword(Password.NewPassword, email)
// 	if err != nil {
// 		r.JSON(400, gin.H{
// 			"message": err.Error})
// 	} else {
// 		r.JSON(200, gin.H{
// 			"message": "password changed"})
// 	}
// }
