package handlers

import (
	"github/jaseelaali/orchid/repository"

	"github.com/gin-gonic/gin"
)

func AddWishList(r *gin.Context) {
	var body struct {
		Product_id int `json:"product_id"`
	}
	err := r.Bind(&body)
	if err != nil {
		r.JSON(400, gin.H{
			"error": "error in binding data",
		})
		return
	}
	user_id := repository.GetId(r)
	err = repository.AddWishlist(user_id, body.Product_id)
	if err != nil {
		r.JSON(400, gin.H{
			"errror": err.Error(),
		})
		return
	}
	r.JSON(200, gin.H{
		"success": "product added to wishlist",
	})
}
func ListWishlist(r *gin.Context) {
	user_id := repository.GetId(r)
	list, err := repository.ViewWishList(user_id)
	if err != nil {
		r.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}
	r.JSON(200, gin.H{
		"WishList": list,
	})
}
func RemoveWishlist(r *gin.Context) {
	var body struct {
		Product_id int `json:"product_id"`
	}
	err := r.Bind(&body)
	if err != nil {
		r.JSON(400, gin.H{
			"error": "error in binding data",
		})
		return
	}
	user_id := repository.GetId(r)
	err = repository.RemoveWishList(user_id, body.Product_id)
	if err != nil {
		r.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}
	r.JSON(400, gin.H{
		"success": "product delete from user wishlist",
	})
}
