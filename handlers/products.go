package handlers

import (
	"fmt"
	"github/jaseelaali/orchid/models"
	"github/jaseelaali/orchid/repository"
	"strconv"

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
	id, err := strconv.Atoi(r.Query("product id"))
	if err != nil {
		r.JSON(400, gin.H{
			"message": "product id didn't get",
		})
		return
	}
	product_name := r.Query("product name")
	product_colour := r.Query("product colour")
	product_size, err := strconv.Atoi(r.Query("product size"))
	product_brand := r.Query("product brand")
	product_price, err := strconv.Atoi(r.Query("product price"))
	stock, err := strconv.Atoi(r.Query("stock"))

	if product_name != "" {
		err := repository.EditProductName(product_name, id)
		if err != nil {
			r.JSON(400, gin.H{
				"message": err.Error(),
			})
			return
		}

	}
	if product_colour != "" {
		err := repository.EditProductColour(product_colour, id)
		if err != nil {
			r.JSON(400, gin.H{
				"message": err.Error(),
			})
			return
		}
	}
	if product_size != 0 {
		err = repository.EditProductSize(product_size, id)
		if err != nil {
			r.JSON(400, gin.H{
				"message": err.Error(),
			})
			return
		}
	}
	if product_brand != "" {
		err = repository.EditProductBrand(product_brand, id)
		if err != nil {
			r.JSON(400, gin.H{
				"message": err.Error(),
			})
			return
		}
	}
	if product_price != 0 {
		err = repository.EditProductPrice(product_price, id)
		if err != nil {
			r.JSON(400, gin.H{
				"message": err.Error(),
			})
			return
		}
	}
	if stock != 0 {
		err = repository.EditProductStock(stock, id)
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
	id, err := strconv.Atoi(r.Query("id"))

	err = repository.Deleteproduct(id)
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
	page, err := strconv.Atoi(r.Query("page"))
	if err != nil {
		r.JSON(400, gin.H{
			"error": "didn't get page number",
		})
		return
	}
	perpage, err := strconv.Atoi(r.Query("perpage"))
	if err != nil {
		r.JSON(400, gin.H{
			"error": "didn't get perpage number",
		})
		return
	}

	Products, metadata, err := repository.Viewproduct(page, perpage)
	if err != nil {

		r.JSON(400, gin.H{
			"error": err.Error})
		return
	}
	r.JSON(200, gin.H{
		"products": Products,
		"metadata": metadata,
	})
}
