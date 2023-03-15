package models

type Cart struct {
	Cart_id     int16 `json:"cart_id" gorm:"unique;primarykey;AUTO_INCREMENT"`
	User_id     int
	Product_id  int
	
	Total_price int
}
type CartResponse struct{

}