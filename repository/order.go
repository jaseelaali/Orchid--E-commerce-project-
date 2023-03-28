package repository

import (
	"fmt"
	"github/jaseelaali/orchid/database"
	"github/jaseelaali/orchid/models"
)

func Add_Order(user_id int) (int, error) {
	//:=models.Cart{}
	//database.DB.Raw("")
	var cart, Address_id int
	result := database.DB.Raw("SELECT sum(product_total_price) FROM cart_items WHERE user_id=$1 ;", user_id).Scan(&cart)
	if result.Error != nil {
		return 0, result.Error
	}
	fmt.Printf("*****%v", cart)

	result = database.DB.Raw("SELECT id FROM addresses WHERE user_id=$1;", user_id).Scan(&Address_id)
	if result.Error != nil {
		return 0, result.Error
	}
	result = database.DB.Raw("INSERT INTO orders (user_id,address_id,total_Amount)VALUES(?,?,?);", user_id, Address_id, cart).Scan(&models.Order{})
	if result.Error != nil {
		return 0, result.Error
	}
	return cart, nil
}
func OrderUpdation(payment_id string, user_id int) error {
	result := database.DB.Raw("UPDATE orders SET payment_status=$1,payment_id=$2,order_status=$3 WHERE user_id=$4;", "success", payment_id, "success", user_id).Scan(&models.Order{})
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func Show_Order(user_id int) ([]models.OrderStatus, error) {
	body := []models.OrderStatus{}
	result := database.DB.Raw("SELECT product_name,quantity,product_price,delivery FROM order_statuses WHERE user_id=$1;", user_id).Scan(&body)
	if result.Error != nil {
		return body, result.Error
	}
	return body, nil
}
