package handlers

import (
	"github/jaseelaali/orchid/models"
	"github/jaseelaali/orchid/repository"

	"github.com/gin-gonic/gin"
)

func Profile(r *gin.Context) {
	user_id := repository.GetId(r)
	profile, err := repository.ViewProfile(user_id)
	if err != nil {
		r.JSON(400, gin.H{
			"message": err.Error(),
		})
		return
	}
	r.JSON(200, gin.H{
		"INFO": profile,
	})

}
func EditProfile(r *gin.Context) {
	user_id := repository.GetId(r)
	body := models.UserProfile{}
	err := r.Bind(&body)
	if err != nil {
		r.JSON(400, gin.H{
			"message": "error in binding data",
		})
		return
	}
	if body.Email != "" {
		err = repository.EditProfileEmail(body.Email, user_id)
	}
	if body.Phone_Number != "" {
		err = repository.EditProfilePhoneNumber(body.Phone_Number, user_id)
	}
	if body.User_Name != "" {
		err = repository.EditProfileUserName(body.User_Name, user_id)
	}
	if err != nil {
		r.JSON(400, gin.H{
			"message": err.Error(),
		})
		return
	}
	r.JSON(200, gin.H{
		"message": "profile updated",
	})
}
func DeleteProfile(r *gin.Context) {
	user_id := repository.GetId(r)
	err := repository.Deleteprofile(user_id)
	if err != nil {
		r.JSON(400, gin.H{
			"message": err.Error(),
		})
		return

	}
	r.JSON(200, gin.H{
		"message": "profile deleted",
	})
}
