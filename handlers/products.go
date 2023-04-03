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
func EditProducts(r *gin.Context) {

	var body struct {
		Id             int    `json:"id" binding:"required"`
		Product_Name   string `json:"product_name"`
		Product_Colour string `json:"product_colour"`
		Product_Size   int    `json:"product_size"`
		Product_Brand  string `json:"product_brand"`
		Product_Price  int    `json:"product_price"`
		Stock          int    `json:"stock"`
	}
	err := r.Bind(&body)
	if err != nil {
		r.JSON(400, gin.H{
			"message": err.Error(),
		})
		return
	}
	if body.Product_Name != "" {
		err := repository.EditProductName(body.Product_Name, body.Id)
		if err != nil {
			r.JSON(400, gin.H{
				"message": err.Error(),
			})
			return
		}

	}
	if body.Product_Size != 0 {
		err = repository.EditProductSize(body.Product_Size, body.Id)
		if err != nil {
			r.JSON(400, gin.H{
				"message": err.Error(),
			})
			return
		}
	}
	if body.Product_Brand != "" {
		err = repository.EditProductBrand(body.Product_Brand, body.Id)
		if err != nil {
			r.JSON(400, gin.H{
				"message": err.Error(),
			})
			return
		}
	}
	if body.Product_Price != 0 {
		err = repository.EditProductPrice(body.Product_Price, body.Id)
		if err != nil {
			r.JSON(400, gin.H{
				"message": err.Error(),
			})
			return
		}
	}
	if body.Stock != 0 {
		err = repository.EditProductStock(body.Stock, body.Id)
		if err != nil {
			r.JSON(400, gin.H{
				"message": err.Error(),
			})
			return
		}
	}
	r.JSON(200, gin.H{
		"message": "product updated",
	})

}
func DeleteProducts(r *gin.Context) {

	var body struct {
		Product_id int `json:"product_id" binding:"required"`
	}
	err := r.Bind(&body)
	if err != nil {
		r.JSON(400, gin.H{
			"message": err.Error(),
		})
		return
	}
	err = repository.Deleteproduct(body.Product_id)
	if err != nil {
		r.JSON(400, gin.H{
			"message": err.Error(),
		})
		return
	}
	r.JSON(200, gin.H{
		"message": "product deleted",
	})

}
func ViewProducts(r *gin.Context) {
	Products, err := repository.Viewproduct()
	if err != nil {

		r.JSON(400, gin.H{
			"error": err.Error})
		return
	}
	r.JSON(200, gin.H{
		"products": Products,
	})
}
