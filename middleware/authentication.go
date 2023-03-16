package middleware

import (
	//"bytes"

	"fmt"
	"github/jaseelaali/orchid/database"
	"github/jaseelaali/orchid/models"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
)

func RequiredAuthentication(r *gin.Context) {
	tokenString, err := r.Cookie("Authorization")

	if err != nil {
		r.JSON(http.StatusUnauthorized, gin.H{"error": "Plese login first"})

	}
	token, _ := jwt.Parse(tokenString, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", t.Header["alg"])

		}
		return []byte(os.Getenv("SECRET")), nil
	})

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		if float64(time.Now().Unix()) > claims["exp"].(float64) {
			r.JSON(http.StatusUnauthorized, gin.H{"error": "Autherization failed"})
			r.Abort()
		}
		user := models.User{}
		result := database.DB.Where("email", claims["sub"]).First(&user)
		if result.Error != nil {
			r.JSON(http.StatusUnauthorized, gin.H{"error": "not user found!"})
		}
		r.Set("user_id", user.ID)
		r.Next()
	} else {
		r.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Autherization failed !"})

	}
}
