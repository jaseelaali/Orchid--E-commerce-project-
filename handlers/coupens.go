package handlers

import (
	"fmt"
	"github/jaseelaali/orchid/repository"
	"time"

	"github.com/gin-gonic/gin"
)

func AddCoupens(r *gin.Context) {

	var body struct {
		Code string `json:"code"`

		MinAmount int `json:"minamount"`
		Amount    int `json:"amount"`
	}
	err := r.Bind(&body)
	if err != nil {
		r.JSON(400, gin.H{
			"error": "error in binding data",
		})
		return
	}
	var expiry time.Time
	expiry = time.Now().Add(time.Hour * 24 * 1)
	err = repository.Addcoupen(body.Code, expiry, body.MinAmount, body.Amount)
	if err != nil {
		r.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}
	r.JSON(200, gin.H{
		"success": "coupen added successfully",
	})
}
func ListCoupen(r *gin.Context) {
	coupen, result := repository.Listcoupen()
	if result != nil {
		r.JSON(400, gin.H{
			"error": result.Error(),
		})
		return
	}
	r.JSON(200, gin.H{
		"coupens": coupen,
	})
}

func ApplyCoupen(r *gin.Context) {
	var body struct {
		Coupenname string `json:"coupenname"`
	}
	err := r.Bind(&body)
	if err != nil {
		r.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}
	user_id := repository.GetId(r)
	fmt.Printf("+++++++:%v", user_id)
	data, err := repository.Applycoupen(user_id, body.Coupenname)
	if err != nil {
		r.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}
	r.JSON(200, gin.H{
		"data":    data,
		"success": "coupen apply succesfully",
	})

}
