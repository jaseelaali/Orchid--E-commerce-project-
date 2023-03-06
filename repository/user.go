package repository

import (
	"github/jaseelaali/orchid/database"
	"github/jaseelaali/orchid/models"
)

func CreateUser(newUser models.User) error {
	user := models.User{}
	result := database.DB.Raw("INSERT INTO users(user_name,email,phone_number,password) VALUES($1,$2,$3,$4);",
		newUser.User_Name, newUser.Email, newUser.Phone_Number, newUser.Password).Scan(&user)
	if result.Error != nil {
		return result.Error
	}
	return nil

}
