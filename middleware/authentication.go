package middleware

/*
import (
	"bytes"
	"fmt"
	"github/jaseelaali/orchid/models"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)
func RequiredAuthentication(r*gin.Context){
//get the cookie of request
tokenstring,err := r.Cookie("Authorization")
if err!=nil{
	r.AbortWithStatus(http.StatusUnauthorized)
}
//decode or validate it
token,_:=jwt.Parse(tokenstring, func(token *jwt.Token) (interface{}, error) {
	if_,ok :=token.Method.(*jwt.SigningMethodHMAC);!ok{
		return nil,fmt.Errorf("unexpected signing method:%v",token.Header["alg"])
	}

return []byte(os.Getenv("SECRET")),nil
})
if claims ,ok:=token.claims.(jwt.MapClaims);ok &&  token.Valid{
	// 	check the exp
	if float64(time.Now().Unix())>claims["exp"].(float64){
      r.AbortWithStatus(http.StatusUnauthorized)
	}
//find the user with token sub
var user models.User
database.DB.first(&user,claims["sub"])}
if User.Id==0{
	r.AbortWithStatus(http.StatusUnauthorized)
}
//attach to request
r.Set("user",user)
//continue
r.Next()

}else{
   c.AbortWithStatus(http.StatusUnauthorized)
}

*/
