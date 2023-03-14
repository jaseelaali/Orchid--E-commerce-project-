package repository

import (
	"github/jaseelaali/orchid/database"
	"github/jaseelaali/orchid/models"
)

func Addcategory(category models.Category) error {
	Category := models.Category{}
	err := database.DB.Raw("INSERT INTO categories(category_name)VALUES ($1)",
		category.Category_Name).Scan(&Category)
	//fmt.Println("*1****** %v", category.Category_Name)
	if err != nil {
		return err.Error
	}
	return nil
}
