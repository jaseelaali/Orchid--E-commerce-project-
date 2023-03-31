package repository

import (
	"errors"
	"fmt"
	"github/jaseelaali/orchid/database"
	"github/jaseelaali/orchid/models"
	"time"
)

func Addcoupen(code string, expiry time.Time, minamount, amount int) error {

	fmt.Printf("----------%v", expiry)
	result := database.DB.Raw("INSERT INTO coupens(code,expiry,min_amount,amount,status)VALUES($1,$2,$3,$4,$5)", code, expiry, minamount, amount, "true").Scan(&models.Coupen{})
	if result.Error != nil {
		return result.Error
	}

	return nil
}
func Listcoupen() ([]models.Coupen, error) {
	err := Updatecoupen()
	if err != nil {
		return nil, err
	}
	body := []models.Coupen{}
	result := database.DB.Raw("SELECT * FROM coupens WHERE status=$1;", "true").Scan(&body)
	if result.Error != nil {
		return body, result.Error
	}
	return body, nil
}
func Updatecoupen() error {
	body := []models.Coupen{}
	//time:=time.Now()
	result := database.DB.Raw("SELECT expiry FROM coupens WHERE status=$1;", "true").Scan(&body)
	if result.Error != nil {
		return result.Error
	}
	for i := range body {
		if body[i].Expiry.Before(time.Now()) {
			result := database.DB.Raw("UPDATE coupens SET status=$1;", "false").Scan(&models.Coupen{})
			if result.Error != nil {
				return result.Error
			}
		}

	}
	return nil

}
func Applycoupen(user_id int, name string) (int, error) {
	body := models.Coupen{}
	body1 := models.Order{}
	result := database.DB.Raw("SELECT * FROM coupens WHERE code=$1;", name).Scan(&body)
	if result.Error != nil {
		return 0, result.Error
	}
	if body.Status == "true" {
		result = database.DB.Raw("SELECT * FROM orders WHERE user_id=$1 and coupen is null and coupen_name is null;", user_id).Scan(&body1)
		//result = database.DB.Where("user_id", user_id).First(&body1)
		if result.Error != nil {
			return 0, result.Error
		}
		var names int
		result = database.DB.Raw("SELECT count(coupen_name) FROM orders WHERE user_id=$1 and coupen_name=$2 and coupen !=$3;", user_id, name, "not-eligible").Scan(&names)
		if result.Error != nil {
			return 0, result.Error
		}
		fmt.Println("count is ", names)
		if names == 0 {
			fmt.Printf("totalamount from cart:%v\n", body1.TotalCartAmount)
			fmt.Printf(":dicount:%v\n", body.Amount)

			if body.MinAmount <= body1.TotalCartAmount {
				result := database.DB.Raw("UPDATE orders SET coupen=$1,coupen_name=$2,total_amount=$3 WHERE user_id=$4;", "applied", name, body1.TotalCartAmount-body.Amount, user_id).Scan(&models.Order{})
				if result.Error != nil {
					return 0, result.Error
				} else {
					Data := body1.TotalCartAmount - body.Amount
					return Data, nil
				}

			} else {
				result := database.DB.Raw("UPDATE orders SET coupen=$1,total_amount=$2 WHERE user_id=$3;", "not-eligible", body1.TotalCartAmount, user_id).Scan(&models.Order{})
				if result.Error != nil {
					return 0, result.Error
				}
				return 0, errors.New("you are not eligible for coupen offer")
			}
		} else {
			database.DB.Raw("UPDATE orders SET total_amount=$1 WHERE user_id=$2;", body1.TotalCartAmount, user_id).Scan(&models.Order{})
			return 0, errors.New("you are already use this coupen")

		}
	} else {
		return 0, errors.New("coupen is not valid")
	}

}
