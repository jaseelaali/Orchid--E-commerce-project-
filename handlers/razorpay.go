package handlers

import (
	"fmt"
	"strings"

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

	total_price := 2000 * 100

	client := razorpay.NewClient("rzp_test_7iVTUnCT2A4xG5", "JAUioUJ7ZkBOcwLmXwN85hQ5")
	razorpaytotal := total_price
	data := map[string]interface{}{
		"amount":   razorpaytotal,
		"currency": "INR",
	}
	body, err := client.Order.Create(data, nil)
	fmt.Println(body)
	if err != nil {
		r.HTML(422, "failed to create order", nil)
	}
	value := fmt.Sprint(body["id"])
	Home := Home{
		userid:      "1",
		Name:        "jaseela",
		total_price: total_price,
		Amount:      total_price,
		OrderId:     value,
		Email:       "jaseelaali2000@gmail.com",
		Contact:     "9909089079",
	}
	// order := domain.Orders{
	// 	Created_at:      time.Now(),
	// 	User_Id:         user.ID,
	// 	Order_Id:        value,
	// 	Applied_Coupons: Coupn.Coupon,
	// 	Discount:        uint(Coupn.Discount),
	// 	Total_Amount:    uint(sum),
	// 	Balance_Amount:  sum - Coupn.Discount,
	// 	PaymentMethod:   "razorpay",
	// 	Payment_Status:  "incomplete",
	// 	Order_Status:    "order_placed",
	// 	Address_Id:      uint(address_id),
	// }
	// err = cr.UserService.CreateOrder(order)
	// if err != nil {
	// 	c.HTML(422, "faile to create order", nil)
	// }

	r.HTML(200, "app.html", Home)

}

func Payment_Success(r *gin.Context) {
	payment_id := r.Query("paymentid")
	orderid := r.Query("orderid")
	orderid = strings.Trim(orderid, " ")
	signature := r.Query("signature")
	fmt.Println(payment_id, signature)
	r.JSON(200, gin.H{
		//"message": "payment success",engine.LoadHTMLGlob("templates/*.html")
	})

}
