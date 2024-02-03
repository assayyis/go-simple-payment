package services

import (
	"assessment/merchant-bank-payment/app/constants"
	"assessment/merchant-bank-payment/app/types"
	"assessment/merchant-bank-payment/app/utils"
	"encoding/json"
	"fmt"
	"log"
	"os"
)

func Pay(customerId int, merchantId int, amount int) string {
	content, err := os.ReadFile(constants.DATA_PATH)
	if err != nil {
		log.Fatal(err)
		return constants.MESSAGE_PAYMENT_FAILED + ", Cause: " + err.Error()
	}
	data := types.Data{}
	err = json.Unmarshal(content, &data)
	if err != nil {
		log.Fatal(err)
		return constants.MESSAGE_PAYMENT_FAILED + ", Cause: " + err.Error()
	}

	customers := data.Customers
	for _, customer := range customers {
		if customer.Id == customerId && !customer.LoggedIn {
			return constants.MESSAGE_PAYMENT_FAILED + ", Cause: " + constants.MESSAGE_LOGIN_TO_PAY
		}
	}

	history := "[" + utils.Now() + "] Payment amount $" + fmt.Sprint(amount) + " succeed with Customer Id " + fmt.Sprint(customerId) + " and Merchant Id " + fmt.Sprint(merchantId)
	data.Histories = append(data.Histories, history)

	content, err = json.Marshal(data)
	if err != nil {
		log.Fatal(err)
		return constants.MESSAGE_PAYMENT_FAILED + ", Cause: " + err.Error()
	}

	err = os.WriteFile(constants.DATA_PATH, content, 0644)
	if err != nil {
		log.Fatal(err)
		return constants.MESSAGE_PAYMENT_FAILED + ", Cause: " + err.Error()
	}

	return constants.MESSAGE_PAYMENT_SUCCESS

}
