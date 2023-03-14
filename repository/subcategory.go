package repository

import (
	"github/jaseelaali/orchid/database"
	"github/jaseelaali/orchid/models"
)

func AddSubcategory(Subcategory models.SubCategory) error {
	SubCategory := models.SubCategory{}
	err := database.DB.Raw("INSERT INTO sub_categories(sub_category_name)VALUES ($1);", Subcategory.SubCategory_Name).Scan(&SubCategory)
	// fmt.Println("******* %v", Subcategory.SubCategory_Name)
	if err != nil {
		return err.Error
	}
	return nil
}
