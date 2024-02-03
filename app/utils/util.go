package utils

import (
	"assessment/merchant-bank-payment/app/types"
	"time"
)

func Contains(customers []types.Customer, username string, password string) bool {
	for _, customer := range customers {
		if customer.Username == username && customer.Password == password {
			return true
		}
	}
	return false
}

func Now() string {
	t := time.Now()
	return t.Format("2006-01-02 15:04:05")
}
