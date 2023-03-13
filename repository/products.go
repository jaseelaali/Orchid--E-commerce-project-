package repository

import (
	"github/jaseelaali/orchid/database"
	"github/jaseelaali/orchid/models"
)

func Addproduct(product models.Products) error {
	newproduct := models.Products{}
	err := database.DB.Raw("INSER INTO produts(product_name,product_colour,product_size,product_brand)VALUES($1,$2,$3,$4);", product.Product_Name,
		product.Product_Colour, product.Product_Size, product.Product_Brand).Scan(&newproduct)
	if err != nil {
		return err.Error
	}
	return nil
}
