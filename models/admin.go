package models

import "gorm.io/gorm"

type Admin struct {
	gorm.Model
	Admin_Name string
	Email      string
	Password   string
}
type User struct {
	gorm.Model
	User_Name    string
	Email        string
	Phone_Number string
	password     string
}
