package models

type Cart struct {
	Cart_id int `json:"cart_id" gorm:"unique;primarykey;AUTO_INCREMENT"`
	User_id int
}
type CartItem struct {
	CartI_Id      int `json:"carti_id" gorm:"unique;primarykey"`
	User_id       int `json:"user_id"`
	Product_Name  string
	Product_Id    int
	Quantity      int
	Product_Price int
	Total_price   int
}
