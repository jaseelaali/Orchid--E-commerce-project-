package handlers

import (
	"github/jaseelaali/orchid/repository"
	"strconv"

	"github.com/gin-gonic/gin"
)

func AddCart(r *gin.Context) {
	var Cart struct {
		Product_Id       int `json:"product_id" binding:"required"`
		Product_Quantity int `json:"product_quantity" binding:"required"`
	}
	err := r.Bind(&Cart)
	if err != nil {
		r.JSON(400, gin.H{
			"message": "error in binding data",
		})
		return
	}

	User_Id, _ := strconv.Atoi(r.Writer.Header().Get("id"))
	err = repository.Addcart(Cart.Product_Id, Cart.Product_Quantity, User_Id)
	if err != nil {
		r.JSON(400, gin.H{
			"message": err.Error(),
		})
		return
	}
	r.JSON(200, gin.H{
		"message": "product added to cart sucessfully",
	})
}
