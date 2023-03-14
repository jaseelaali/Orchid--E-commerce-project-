package handlers

import (
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
