package repository

import (

	"github/jaseelaali/orchid/database"
	"github/jaseelaali/orchid/models"
)

func Addproduct(product models.Products) error {
	err := database.DB.Create(&models.Products{
		Product_Name:   product.Product_Name,
		Product_Colour: product.Product_Colour,
		Product_Size:   product.Product_Size,
		Product_Brand:  product.Product_Brand,
		Product_Price:  product.Product_Price,
	})
	// fmt.Println("***********************")
	// fmt.Println(product.Product_Size)
	// fmt.Println("****************")
	// // fmt.Println()
	// fmt.Println()
	// fmt.Println()

	// fmt.Println()

	// err := database.DB.Raw("INSERT INTO products (product_name,product_colour,product_size,product_brand,product_price) VALUES ($1,$2,$3,$4,$5);", product.Product_Name,
	// 	product.Product_Colour, product.Product_Size, product.Product_Brand, product.Product_Price).Scan(&newproduct)
	if err.Error != nil {
		return err.Error
	}
	return nil
}
