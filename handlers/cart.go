package handlers

import (
	"fmt"
	"github/jaseelaali/orchid/database"
	"github/jaseelaali/orchid/repository"
	"strconv"

	"github.com/gin-gonic/gin"
	//"github.com/go-playground/validator/v10/translations/id"
)

func AddCart(r *gin.Context) {
	var body struct {
		Product_Id       int `json:"product_id" binding:"required"`
		Product_Quantity int `json:"product_quantity" binding:"required"`
	}

	user_id := repository.GetId(r)
	err := r.Bind(&body)
	if err != nil {
		r.JSON(400, gin.H{
			"message": "error in binding data",
		})
		return
	}
	var stock int
	err = database.DB.Raw("SELECT stock FROM products WHERE id=$1;", body.Product_Id).Scan(&stock).Error
	if err != nil {
		r.JSON(400, gin.H{
			"message": err.Error(),
		})
		return
	}
	if body.Product_Quantity > stock {
		r.JSON(400, gin.H{
			"message": "the product is out of stock",
		})
		return
	}
	var alreadyexist int
	database.DB.Raw("SELECT quantity FROM cart_items WHERE product_id=$1 AND user_id=$2;", body.Product_Id, user_id).Scan(&alreadyexist)
	if alreadyexist != 0 {
		err = repository.ADDcart(alreadyexist, body.Product_Id, body.Product_Quantity, user_id)
		if err != nil {
			r.JSON(400, gin.H{
				"message": err.Error(),
			})
			return
		}
		sum, err := repository.SumCart(user_id)
		if err != nil {
			r.JSON(400, gin.H{
				"message": err.Error(),
			})
			return
		}
		r.JSON(200, gin.H{
			"sum of the products is": sum,
		})

	} else {
		//database.DB.Raw("UPDATE PRODUCTS SET stock=$1 WHERE id=$2;", stock-body.Product_Quantity, body.Product_Id).Scan(&models.Products{})
		err = repository.Addcart(body.Product_Id, body.Product_Quantity, user_id)
		if err != nil {
			r.JSON(400, gin.H{
				"message": err.Error(),
			})
			return
		}
		sum, err := repository.SumCart(user_id)
		if err != nil {
			r.JSON(400, gin.H{
				"message": err.Error(),
			})
			return
		}
		r.JSON(200, gin.H{
			"sum of the products is": sum,
		})

	}

	r.JSON(200, gin.H{

		"message": "product added to cart successfully",
	})
}

func ViewCart(r *gin.Context) {
	var Body struct {
		Page    int `json:"page" binding:"required"`
		Perpage int `json:"perpage" binding:"required"`
	}
	err := r.ShouldBind(&Body)
	if err != nil {
		r.JSON(400, gin.H{
			"message": err.Error()})
		return
	}
	user_id, _ := r.Get("user_id")
	userID, _ := strconv.Atoi((fmt.Sprint(user_id)))

	cart, metadata, err := repository.Viewcart(userID, Body.Page, Body.Perpage)
	if err != nil {
		r.JSON(400, gin.H{
			"message": err.Error(),
		})
		return
	}

	r.JSON(200, gin.H{
		"cart":     cart,
		"metadata": metadata})

}
func DeleteItem(r *gin.Context) {
	var body struct {
		Id       int `json:"id" `
		Quantity int `json:"quantity"`
	}
	err := r.Bind(&body)
	if err != nil {
		r.JSON(400, gin.H{
			"message ": "error in binding data",
		})
		return
	}
	user_id, _ := r.Get("user_id")
	userID, _ := strconv.Atoi((fmt.Sprint(user_id)))
	err = repository.Deleteitem(body.Id, body.Quantity, userID)
	//mt.Println("id:%v**quantity:%v**id:%v;", body.Id, body.Quantity, userID)
	if err != nil {
		r.JSON(400, gin.H{
			"message": err.Error(),
		})
		return
	}
	r.JSON(200, gin.H{
		"message": "product delete from cart  successfully",
	})
}
