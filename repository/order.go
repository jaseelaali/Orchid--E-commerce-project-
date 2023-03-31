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
	result = database.DB.Raw("INSERT INTO orders (user_id,address_id,total_cart_amount)VALUES(?,?,?);", user_id, Address_id, cart).Scan(&models.Order{})
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

func Show_Order(user_id int) ([]models.CartItem, error) {
	value := []models.CartItem{}
	result := database.DB.Raw("SELECT * FROM orders WHERE user_id=$1 AND order_status is null;", user_id).Scan(&value)
	if result.Error != nil {
		return nil, result.Error
	}
	if value != nil {

		body := []models.CartItem{}
		//result := database.DB.Raw("SELECT * FROM order_statuses WHERE user_id=$1 AND delivery= 'not done' ;", user_id).Scan(&body)
		result := database.DB.Raw("SELECT product_name,quantity,product_price,product_total_price from cart_items WHERE user_id=$1", user_id).Scan(&body)
		if result.Error != nil {
			return body, result.Error
		}
		return body, nil
	} else {
		return nil, nil
	}
}

func Cancel_Order(user_id int) error {
	fmt.Printf("+++++++++++++++++++++++++++:%v:userid", user_id)

	result := database.DB.Raw("UPDATE orders SET order_status='cancelled' WHERE user_id=$1", user_id).Scan(&models.Order{})
	if result.Error != nil {
		return result.Error
	}
	result = database.DB.Raw("DELETE FROM cart_items WHERE user_id=$1", user_id).Scan(&models.CartItem{})
	if result.Error != nil {
		return result.Error
	}
	return nil
}
