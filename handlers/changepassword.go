package handlers

import (
	// "fmt"
	// "math/rand"
	// "strconv"

	"fmt"
	"github/jaseelaali/orchid/database"
	"github/jaseelaali/orchid/models"
	"math/rand"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/sfreiberg/gotwilio"
	"golang.org/x/crypto/bcrypt"
)

var OTP int
var twilio = gotwilio.NewTwilioClient("AC3037e122f46a35ae97b5a48f7413be56", "0bc43f0b4e4a492d46e26bf093c0fc40")

func ChangePassword(r *gin.Context) {
	user_id, _ := r.Get("user_id")
	User_Id, _ := strconv.Atoi(fmt.Sprint(user_id))
	var mobilenumber string
	err := database.DB.Raw("SELECT phone_number FROM users WHERE id=$1;", User_Id).Scan(&mobilenumber)
	if err.Error != nil {
		r.JSON(400, gin.H{
			"message": err.Error,
		})
		return
	}
	otp, Error := sendOTP(mobilenumber)
	OTP = otp
	if Error != nil {
		r.JSON(400, gin.H{
			"message": "failed to send message",
			"error":   Error,
		})
		return
	}
	r.JSON(200, gin.H{
		"message": "successfully send the otp",
		"data":    OTP,
	})

}
func sendOTP(phoneNumber string) (int, error) {
	otpCode := generateOTP()
	message := "Your OTP code is " + strconv.Itoa(otpCode)
	_, _, err := twilio.SendSMS("+15302036484", "+91"+phoneNumber, message, "", "")
	if err != nil {
		return -1, err
	}
	return otpCode, nil
}
func generateOTP() int {
	// Generate a random 4-digit OTP code
	otp := rand.Intn(8999) + 1000
	fmt.Println(otp)
	return otp
}
func VerifyOtp(r *gin.Context) {
	var body struct {
		Otp          int    `json:"otp" binding:"required"`
		New_Password string `json:"new_password" binding:"required"`
	}
	err := r.Bind(&body)
	if err != nil {
		r.JSON(400, gin.H{
			"message": "error in binding data",
		})
	}
	isValid := VerifyOTP(body.Otp, OTP)
	if isValid == true {
		newpasword := body.New_Password
		password, err := bcrypt.GenerateFromPassword([]byte(newpasword), 11)
		userid, _ := strconv.Atoi(fmt.Sprint(r.Get("user_id")))
		result := database.DB.Raw("UPDATE users SET password=$1 WHERE id=$2", password, userid).Scan(&models.User{})
		if result.Error != nil {
			r.JSON(400, gin.H{
				"message": err.Error(),
			})
			return
		}
		r.JSON(200, gin.H{
			"message": "password changed successfully",
		})
		return
	} else {
		r.JSON(400, gin.H{
			"message": "invalid otp",
		})

	}
}
func VerifyOTP(otpCode, expectedCode int) bool {
	if otpCode == expectedCode {
		return true
	}
	return false
}
