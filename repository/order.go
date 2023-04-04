package repository

import (
	"errors"
	"fmt"
	"github/jaseelaali/orchid/database"
	"github/jaseelaali/orchid/models"
)

func Add_Order(user_id, Address_id int) (int, error) {
	//cart := models.Cart{}
	database.DB.Raw("")
	var cart int
	fmt.Printf("----add----:%v", user_id)
	fmt.Printf("----add----:%v", Address_id)
	var payment string
	result := database.DB.Raw("SELECT payment FROM ordered_products WHERE user_id=$1;", user_id).Scan(&payment)
	if result.Error != nil {
		return 0, result.Error
	}
	if payment == "not done" {

		return 0, errors.New("one order is  pending")
	}

	result = database.DB.Raw("SELECT sum(product_total_price) FROM cart_items WHERE user_id=$1 ;", user_id).Scan(&cart)
	if result.Error != nil {
		return 0, result.Error
	}
	fmt.Printf("*****%v", cart)
	//updated
	var order_id int
	result = database.DB.Raw("INSERT INTO orders (user_id,address_id,total_cart_amount,total_amount)VALUES(?,?,?,?);", user_id, Address_id, cart, cart).Scan(&models.Order{})
	result = database.DB.Raw(" id FROM orders WHERE user_id=$1 AND order_status is null;").Scan(&order_id)
	fmt.Printf("......order..........%v", order_id)
	//result = database.DB.Raw("INSERT INTO orders (total_cart_amount)VALUES(?) WHERE user_id=$2 AND address_id=$3;", cart, user_id, Address_id).Scan(&models.Order{})
	//result = database.DB.Raw("UPDATE Orders SET total_cart_amount=$1 WHERE user_id=$2 AND address_id=$3;", cart, user_id, Address_id).Scan(&models.Order{})
	if result.Error != nil {
		return 0, result.Error
	}
	return cart, nil

}
func OrderUpdation(payment_id string, user_id int) error {
	result := database.DB.Raw("UPDATE orders SET payment_method=$1,payment_status=$2,payment_id=$3,order_status=$4 WHERE user_id=$5 AND order_status is null;", "Razor_pay", "success", payment_id, "success", user_id).Scan(&models.Order{})
	if result.Error != nil {
		return result.Error
	}

	result = database.DB.Raw("UPDATE ordered_products SET payment='done' WHERE user_id=$1 AND payment =$2;", user_id, "not done").Scan(&models.OrderedProduct{})
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func Show_Order(user_id int) ([]models.OrderedProduct, error) {
	value := []models.OrderedProduct{}
	result := database.DB.Raw("SELECT * FROM ordered_products WHERE user_id=$1 and payment=$2;", user_id, "not done").Scan(&value)
	if result.Error != nil {
		return nil, result.Error
	}
	return value, nil
}

func Cancel_Order(user_id int) error {
	result := database.DB.Raw("DELETE from ordered_products WHERE user_id=$1 and payment=$2;", user_id, "not done").Scan(&models.OrderedProduct{})
	if result.Error != nil {
		return result.Error
	}
	return nil
}
