package repository

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
)

func GetId(r *gin.Context)int{
	temp := fmt.Sprint(r.Get("user_id"))
	id := strings.Split(temp, " ")
	Id, _ := strconv.Atoi(id[0])
	return Id
}
