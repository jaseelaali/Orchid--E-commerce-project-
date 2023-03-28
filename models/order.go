package models

import "gorm.io/gorm"

type Order struct {
	gorm.Model
	User_Id int
	//Order_Id       int
	Address_Id     int
	Total_Amount   int
	Payment_Method int
	Payment_Status string
	Payment_Id     string
	Order_Status   string
}
type OrderStatus struct {
	gorm.Model
	User_id       int
	Product_id    int
	order_id      int
	Product_name  string
	Quantity      int
	Product_Price int
	Delivery      string
}
