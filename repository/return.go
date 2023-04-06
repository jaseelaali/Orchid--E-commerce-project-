package repository

import (
	"github/jaseelaali/orchid/database"
	"github/jaseelaali/orchid/models"
)

func ReturnStatusChange(id string) error {
	result := database.DB.Raw("UPDATE FROM order_statuses SET delivery=$1 WHERE payment_id=$2;", "completed", id).Scan(&models.OrderStatus{})
	if result.Error != nil {
		return result.Error
	}
	return nil
}
func ReturnProduct(user_id int, order_id, product_id string) error {
	var order models.Order
	var orderstatus models.OrderStatus
	database.DB.Raw("SELECT * FROM orders WHERE id=$1;", order_id).Scan(&order)
	paymentid := order.Payment_Id
	database.DB.Raw("SELECT *FROM order_statuses WHERE payment_id=$1 AND product_id=$2; ", paymentid, product_id).Scan(&orderstatus)
	total_product_price := orderstatus.Product_Price * orderstatus.Quantity
	if order.Coupen == "applied" {
		coupenname := order.Coupen_Name
		var coupen models.Coupen
		database.DB.Raw("SELECT * FROM coupens WHERE code=$1;", coupenname).Scan(&coupen)
		if coupen.MinAmount <= (order.TotalCartAmount - total_product_price) {
			amount := order.TotalCartAmount - total_product_price
			database.DB.Raw("UPDATE orders SET total_cart_amount=$1,total_amount=$2 WHERE payment_id=$3;", amount, amount-coupen.Amount, paymentid).Scan(&models.Order{})
			database.DB.Raw("UPDATE order_statuses SET delivery=$1 WHERE payment_id=$2 AND product_id=$3;", "returned", paymentid, product_id).Scan(&models.OrderStatus{})
			database.DB.Raw("INSERT INTO wallets(user_id, money)VALUES($1,$2);", user_id, total_product_price).Scan(&models.Wallet{})
		} else {
			price := order.TotalCartAmount - total_product_price
			database.DB.Raw("UPDATE orders SET total_cart_amount=$1,total_amount=$2 WHERE payment_id=$3;", price, price, paymentid).Scan(&models.Order{})
			database.DB.Raw("UPDATE order_statuses SET delivery=$1 WHERE payment_id=$2 AND product_id=$3;", "returned", paymentid, product_id).Scan(&models.OrderStatus{})
			database.DB.Raw("INSERT INTO wallets(user_id, money)VALUES($1,$2);", user_id, total_product_price-coupen.Amount).Scan(&models.Wallet{})

		}
	} else {
		price := order.TotalCartAmount - total_product_price
		database.DB.Raw("UPDATE orders SET total_cart_amount=$1,total_amount=$2 WHERE payment_id=$3;", price, price, paymentid).Scan(&models.Order{})
		database.DB.Raw("UPDATE order_statuses SET delivery=$1 WHERE payment_id=$2 AND product_id=$3;", "returned", paymentid, product_id).Scan(&models.OrderStatus{})
		database.DB.Raw("INSERT INTO wallets(user_id, money)VALUES($1,$2);", user_id, total_product_price).Scan(&models.Wallet{})
	}
	return nil
}
