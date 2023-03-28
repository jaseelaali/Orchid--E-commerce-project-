package handlers

import (
	"fmt"
	"github/jaseelaali/orchid/repository"

	"github.com/gin-gonic/gin"
)

func AddOrder(r *gin.Context) {
	user_id := repository.GetId(r)
	data, err := repository.Add_Order(user_id)
	if err != nil {
		r.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	data1 := "your total price is:" + fmt.Sprint(data)
	r.JSON(200, gin.H{
		"data":    data1,
		"success": "placed order successfully",
	})
}
func ShowOrder(r *gin.Context) {
	user_id := repository.GetId(r)
	order, err := repository.Show_Order(user_id)
	if err != nil {
		r.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}
	r.JSON(200, gin.H{
		"success": order,
	})
}
