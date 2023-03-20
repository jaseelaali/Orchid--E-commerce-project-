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
func Editcategory(name string, Id int) error {

	err := database.DB.Raw("UPDATE categories SET category_name=$1 WHERE id=$2;", name, Id).Scan(&models.Category{})
	if err != nil {
		return err.Error
	}
	return nil
}
func Deletecategory(Id int) error {
	err := database.DB.Raw("DELETE FROM categories WHERE id=$1", Id).Scan(&models.Category{})
	if err != nil {
		return err.Error
	}
	return nil
}
func Viewcategory() ([]models.Category, error) {
	category := []models.Category{}
	err := database.DB.Raw("SELECT * FROM categories;").Scan(&category)
	if err != nil {
		return category, err.Error
	}
	return category, nil
}
func EditSubcategory(name string, Id int) error {

	err := database.DB.Raw("UPDATE sub_categories SET sub_category_name=$1 WHERE id=$2;", name, Id).Scan(&models.SubCategory{})
	if err != nil {
		return err.Error
	}
	return nil
}
func DeleteSubcategory(Id int) error {
	err := database.DB.Raw("DELETE FROM sub_categories WHERE id=$1", Id).Scan(&models.SubCategory{})
	if err != nil {
		return err.Error
	}
	return nil
}
func ViewSubcategory() ([]models.SubCategory, error) {
	subcategory := []models.SubCategory{}
	err := database.DB.Raw("SELECT * FROM sub_categories;").Scan(&subcategory)
	if err != nil {
		return subcategory, err.Error
	}
	return subcategory, nil
}
