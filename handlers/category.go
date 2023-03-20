package handlers

import (
	"fmt"
	"github/jaseelaali/orchid/models"
	"github/jaseelaali/orchid/repository"

	"github.com/gin-gonic/gin"
)

func AddCategory(r *gin.Context) {
	category := models.Category{}
	err := r.Bind(&category)
	if err != nil {
		r.JSON(400, gin.H{
			"message": "error in binding data",
		})
		return
	}
	err = repository.Addcategory(category)
	if err != nil {
		r.JSON(400, gin.H{
			"message": err.Error(),
		})
		return
	}
	r.JSON(200, gin.H{
		"message": "category added successfully",
	})
}
func EditCategory(r *gin.Context) {
	var body struct {
		Name string `json:"name"`
		Id   int    `json:"id"`
	}
	err := r.Bind(&body)
	if err != nil {
		r.JSON(400, gin.H{
			"message": "error in binding data",
		})
		return
	}

	err = repository.Editcategory(body.Name, body.Id)
	if err != nil {
		r.JSON(400, gin.H{
			"message": err.Error(),
		})
		return
	}
	r.JSON(200, gin.H{
		"message": "category updated",
	})

}
func DeleteCategory(r *gin.Context) {
	var body struct {
		Category_id int `json:"category_id"`
	}
	err := r.Bind(&body)
	fmt.Println(body.Category_id)
	if err != nil {
		r.JSON(400, gin.H{
			"message": "error in binding data",
		})
		return
	}
	err = repository.Deletecategory(body.Category_id)
	if err != nil {
		r.JSON(400, gin.H{
			"message": err.Error,
		})
		return
	}
	r.JSON(200, gin.H{
		"message": "category deleted",
	})
}
func ViewCategory(r *gin.Context) {
	Category, err := repository.Viewcategory()
	if err != nil {
		r.JSON(400, gin.H{
			"message": "failed to list category",
		})
		return
	}
	r.JSON(400, gin.H{
		"message": Category,
	})

}
func EditSubCategory(r *gin.Context) {
	var body struct {
		Name string `json:"name"`
		Id   int    `json:"id"`
	}
	err := r.Bind(&body)
	if err != nil {
		r.JSON(400, gin.H{
			"message": "error in binding data",
		})
		return
	}

	err = repository.EditSubcategory(body.Name, body.Id)
	if err != nil {
		r.JSON(400, gin.H{
			"message": err.Error(),
		})
		return
	}
	r.JSON(200, gin.H{
		"message": "subcategory updated",
	})

}
func DeleteSubCategory(r *gin.Context) {
	var body struct {
		SubCategory_id int `json:"subcategory_id"`
	}
	err := r.Bind(&body)
	fmt.Println(body.SubCategory_id)
	if err != nil {
		r.JSON(400, gin.H{
			"message": "error in binding data",
		})
		return
	}
	err = repository.DeleteSubcategory(body.SubCategory_id)
	if err != nil {
		r.JSON(400, gin.H{
			"message": err.Error,
		})
		return
	}
	r.JSON(200, gin.H{
		"message": "subcategory deleted",
	})
}
func ViewSubCategory(r *gin.Context) {
	SubCategory, err := repository.ViewSubcategory()
	if err != nil {
		r.JSON(400, gin.H{
			"message": "failed to list subcategory",
		})
		return
	}
	r.JSON(400, gin.H{
		"message": SubCategory,
	})

}
