package handlers

import (
	"github.com/gin-gonic/gin"
)

func RetunMyProduct(r *gin.Context) {
	var body struct {
		Product_id int    `json:"product_id"`
		Reason     string `json:"reason"`
	}
	err := r.Bind(&body)
	if err != nil {
		r.JSON(400, gin.H{
			"error": "error in binding data",
		})
		return
	}
	// user_id:=repository.GetId(r)
	// err=repository.ReturnProduct(user_id,body.Product_id,body.Reason)

}
