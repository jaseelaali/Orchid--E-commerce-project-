package repository

import (
	"github/jaseelaali/orchid/database"
	"github/jaseelaali/orchid/models"
)

func Addcart(Id, Quantity, UId int) error {
	Product_Id := Id
	Product_Quantity := Quantity
	//User_Id  :=UId
	var Product_Price int
	err := database.DB.Raw("SELECT product_price FROM products WHERE id=$1 ;", Product_Id).Scan(&Product_Price)
	total := (Product_Price * Product_Quantity)
	err = database.DB.Raw("INSERT INTO carts (total_price)VALUES($1);", total).Scan(&models.Cart{})
	if err != nil {
		return err.Error
	}

	return nil
}
