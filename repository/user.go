package repository

import (
	"errors"
	"fmt"
	"github/jaseelaali/orchid/database"
	"github/jaseelaali/orchid/models"

	"golang.org/x/crypto/bcrypt"
)

func CreateUser(newUser models.User) error {
	user := models.User{}
	result := database.DB.Raw("INSERT INTO users(user_name,email,phone_number,password,status) VALUES($1,$2,$3,$4,$5);",
		newUser.User_Name, newUser.Email, newUser.Phone_Number, newUser.Password, newUser.Status).Scan(&user)
	if result.Error != nil {
		return result.Error
	}
	return nil

}
func View() ([]models.UserResponses, error) {
	users := []models.UserResponses{}
	err := database.DB.Raw("SELECT * FROM users;").Scan(&users)
	if err != nil {
		return users, err.Error
	}
	return users, nil
}
func BlockUser(user_id int) error {
	//user_id = user_id
	var status string
	database.DB.Raw("SELECT status FROM users WHERE id=$1;", user_id).Scan(&status)

	fmt.Println(user_id)

	if status == "blocked" {
		return errors.New("Selected user is already blocked")
	}
	err := database.DB.Raw("UPDATE users SET status=$1 WHERE id=$2;", "blocked", user_id).Scan(&models.User{})
	if err.Error != nil {
		return err.Error
	}
	return nil
}
func UnBlockUser(user_id int) error {
	var status string
	database.DB.Raw("SELECT status FROM users WHERE id=$1;", user_id).Scan(&status)
	if status == "active" {
		return errors.New("selected user already unblocked")
	}
	err := database.DB.Raw("UPDATE users SET status=$1 WHERE id=$2", "active", user_id).Scan(&models.User{})
	if err.Error != nil {
		return err.Error
	}
	return nil
}
func BlocUsers() ([]models.UserResponses, error) {
	users := []models.UserResponses{}
	err := database.DB.Raw("SELECT * FROM users WHERE status='blocked';").Scan(&users)
	if err != nil {
		return users, err.Error
	}
	return users, nil
}
func ActiveUser() ([]models.UserResponses, error) {
	users := []models.UserResponses{}
	err := database.DB.Raw("SELECT * FROM users WHERE status='active';").Scan(&users)
	if err != nil {
		return users, err.Error
	}
	return users, nil
}
func NewPassword(password string, email string) error {
	newpassword := password
	user := email
	fmt.Println("++++++")
	fmt.Println(password)
	fmt.Println("+++++")
	//fmt.Println(Id)
	fmt.Println("++++++")
	Password, err := (bcrypt.GenerateFromPassword([]byte(newpassword), 11))
	database.DB.Raw("UPDATE password SET password=$1 WHERE id=$2 ", Password, user).Scan(&models.User{})
	if err != nil {
		return err
	}
	return nil

}
