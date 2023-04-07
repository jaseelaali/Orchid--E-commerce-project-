package handlers

import (
	"github/jaseelaali/orchid/repository"
	"strconv"

	"github.com/gin-gonic/gin"
)

func MyWallet(r *gin.Context) {
	user_id := repository.GetId(r)
	data, err := repository.Mywallet(user_id)
	if err != nil {
		r.JSON(400, gin.H{
			"message": err.Error(),
		})
		return
	}
	r.JSON(200, gin.H{
		"message": "your balance is " + strconv.Itoa(data),
	})
}
