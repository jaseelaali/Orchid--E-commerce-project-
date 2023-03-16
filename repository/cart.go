package repository

import (
	//"fmt"
	"github/jaseelaali/orchid/database"
	"github/jaseelaali/orchid/models"
)

func Addcart(P_Id, Quantity, U_Id int) error {
	var productname string
	var productprice int
	err := database.DB.Raw("SELECT product_name FROM products WHERE id=$1;", P_Id).Scan(&productname)
	err = database.DB.Raw("SELECT product_price FROM products WHERE id=$1;", P_Id).Scan(&productprice)
	//fmt.Println("************* %v"productname)

	err = database.DB.Create(&models.CartItem{
		User_id:       U_Id,
		Product_Name:  productname,
		Product_Id:    P_Id,
		Quantity:      Quantity,
		Product_Price: productprice,
		Total_price:   (productprice * Quantity),
	}).Where("user_id", U_Id)
	if err != nil {
		return err.Error
	}
	return nil
}
func ADDcart(newquantity, P_Id, Quantity, U_Id int) error {
	var productname string
	var productprice int
	err := database.DB.Raw("SELECT product_name FROM products WHERE id=$1;", P_Id).Scan(&productname)
	err = database.DB.Raw("SELECT product_price FROM products WHERE id=$1;", P_Id).Scan(&productprice)
	//fmt.Println("************* %v"productname)
	var quantity int
	quantity = (newquantity + Quantity)
	price := ((newquantity + Quantity) * productprice)
	//fmt.Println("****pazhe:%v****puth:%v***quq:%v", newquantity, Quantity, quantity)
	err = database.DB.Raw("UPDATE cart_items SET quantity=$1 WHERE user_id=$2 AND product_id=$3;", quantity, U_Id, P_Id).Scan(&models.CartItem{})
	err = database.DB.Raw("UPDATE cart_items SET total_price=$1 WHERE product_id=$2  AND user_id=$3;", price, P_Id, U_Id).Scan(&models.CartItem{})
	if err != nil {
		return err.Error
	}
	return nil
}
func Viewcart(ID int) ([]models.CartItem, error) {
	Id := ID
	//fmt.Println("************* %v", Id)
	cart := []models.CartItem{}
	err := database.DB.Raw("SELECT * FROM cart_items WHERE  user_id=$1 ;", Id).Scan(&cart)
	if err.Error != nil {
		return cart, err.Error
	}

	return cart, nil

}
func Deleteitem(id, QUantity, User_Id int) error {
	Id := id
	Quantity := QUantity
	U_Id := User_Id
	var OldQuantity, Price int
	err := database.DB.Raw("SELECT quantity FROM cart_items WHERE user_id=$1 AND product_id=$2;", U_Id, Id).Scan(&OldQuantity)
	err = database.DB.Raw("SELECT product_price FROM cart_items WHERE user_id=$1 AND product_id=$2;", U_Id, Id).Scan(&Price)

	quantity := (OldQuantity - Quantity)
	price := Price * quantity
	//fmt.Println("****old:%v**kal:%v**new:%v***price:%v", OldQuantity, Quantity, quantity, price)
	err = database.DB.Raw("UPDATE cart_items SET quantity=$1 WHERE user_id=$2 AND product_id=$3;", quantity, U_Id, Id).Scan(&models.CartItem{})
	err = database.DB.Raw("UPDATE cart_items SET total_price=$1 WHERE user_id=$2 AND product_id=$3;", price, U_Id, Id).Scan(&models.CartItem{})

	if quantity == 0 {
		database.DB.Raw("DELETE FROM cart_items WHERE user_id=$1 AND product_id=$2;", U_Id, Id).Scan(&models.CartItem{})
	}
	if err.Error != nil {
		return err.Error
	}
	return nil
}
