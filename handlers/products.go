package handlers

import (
	"fmt"
	"github/jaseelaali/orchid/models"
	"github/jaseelaali/orchid/repository"

	"github.com/gin-gonic/gin"
)

func AddProducts(r *gin.Context) {
	product := models.Products{}
	err := r.Bind(&product)
	fmt.Printf("\nname : %v\ncolor : %v\nsize : %v\nbrand : %v\nerror : %v\n", product.Product_Name, product.Product_Colour, product.Product_Size, product.Product_Brand, err)
	if err != nil {
		r.JSON(400, gin.H{
			"message": "error in binding data",
		})
		return
	}
	err = repository.Addproduct(product)
	if err != nil {
		r.JSON(400, gin.H{
			"message": err.Error(),
		})
		return
	}
	r.JSON(200, gin.H{"message": " add product successfully"})
}
