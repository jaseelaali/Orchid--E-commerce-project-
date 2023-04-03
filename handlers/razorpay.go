package handlers

import (
	"fmt"
	"github/jaseelaali/orchid/models"
	"github/jaseelaali/orchid/repository"
	"strconv"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	razorpay "github.com/razorpay/razorpay-go"
)

type Home struct {
	userid      string
	Name        string
	total_price int
	Amount      int
	OrderId     string
	Email       string
	Contact     string
}

func Razorpay(r *gin.Context) {
	user_id := 24
	address_id := 19
	fmt.Println(user_id)

	total_price, err := repository.Sum(user_id)
	if err != nil {
		fmt.Println("Couldn't found the total price")
	}
	fmt.Printf("total price is %v", total_price)
	client := razorpay.NewClient("rzp_test_7iVTUnCT2A4xG5", "JAUioUJ7ZkBOcwLmXwN85hQ5")
	razorpaytotal := total_price * 100
	data := map[string]interface{}{
		"amount":   razorpaytotal,
		"currency": "INR",
	}
	body, err := client.Order.Create(data, nil)
	fmt.Println(body)
	if err != nil {
		fmt.Println("Something went wrong")
		r.HTML(422, "failed to create order", nil)
	}
	Order_Id := fmt.Sprint(body["id"])
	order_id, _ := strconv.Atoi(Order_Id)
	Home := Home{
		userid:      "24",
		Name:        "jaseela",
		total_price: total_price,
		Amount:      total_price,
		OrderId:     Order_Id,
		Email:       "jaseela@gmail.com",
		Contact:     "9909089079",
	}
	//value:=repository.Value(user_id)
	//address_id := repository.Address(user_id)
	Payment := models.Payment{
		Created_at: time.Now(),
		User_Id:    user_id,
		Order_Id:   order_id,
		//Applied_Coupons: Coupn.Coupon,
		//Discount:        int(Coupn.Discount),
		//	Total_Amount: int(sum),
		//Balance_Amount:  sum - Coupn.Discount,
		PaymentMethod:  "razorpay",
		Payment_Status: "incomplete",
		Order_Status:   "order_placed",
		Address_Id:     address_id,
	}
	//err = r.UserService.CreateOrder(paymen)
	fmt.Println(Payment)
	if err != nil {
		r.HTML(422, "faile to create order", nil)
	}

	r.HTML(200, "app.html", Home)

}

func Payment_Success(r *gin.Context) {
	r.HTML(200, "success.html", nil)
	payment_id := r.Query("paymentid")
	orderid := r.Query("orderid")
	orderid = strings.Trim(orderid, " ")
	fmt.Println("hello jaseela...")
	// err := repository.RazorPay(payment_id, orderid, Home)
	// if err != nil {
	// 	r.JSON(400, gin.H{
	// 		"message": err,
	// 	})
	// 	return
	// }
	signature := r.Query("signature")
	user_id := 24
	err := repository.OrderUpdation(payment_id, user_id)
	if err != nil {
		r.JSON(400, gin.H{
			"message": err.Error(),
		})
		return
	}
	err = repository.OrderStatus(user_id)
	if err != nil {
		r.JSON(400, gin.H{
			"message": err.Error(),
		})
		return
	}
	// err = repository.ClearCart(user_id)
	// if err != nil {
	// 	r.JSON(400, gin.H{
	// 		"message": err.Error(),
	// 	})
	// 	return
	// }
	fmt.Println(payment_id, signature)
	r.JSON(200, gin.H{
		"message": "payment success",
	})

}
