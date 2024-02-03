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

func Login(username string, password string) string {
	content, err := os.ReadFile(constants.DATA_PATH)
	if err != nil {
		log.Fatal(err)
		return constants.MESSAGE_LOGIN_FAILED + ", Cause: " + err.Error()
	}
	data := types.Data{}
	err = json.Unmarshal(content, &data)
	if err != nil {
		log.Fatal(err)
		return constants.MESSAGE_LOGIN_FAILED + ", Cause: " + err.Error()
	}

	if !utils.Contains(data.Customers, username, password) {
		return constants.MESSAGE_LOGIN_FAILED + ", Cause: " + constants.MESSAGE_INVALID_CREDENTIAL
	}

	customers := data.Customers
	for i, customer := range customers {
		if customer.Username == username {
			customers[i].LoggedIn = true
			break
		}
	}
	data.Customers = customers

	history := "[" + utils.Now() + "] Customer with username " + username + " logged in"
	data.Histories = append(data.Histories, history)

	content, err = json.Marshal(data)
	if err != nil {
		log.Fatal(err)
		return constants.MESSAGE_LOGIN_FAILED + ", Cause: " + err.Error()
	}

	err = os.WriteFile(constants.DATA_PATH, content, 0644)
	if err != nil {
		log.Fatal(err)
		return constants.MESSAGE_LOGIN_FAILED + ", Cause: " + err.Error()
	}

	return constants.MESSAGE_LOGIN_SUCCESS
}

func Logout(id int) string {
	content, err := os.ReadFile(constants.DATA_PATH)
	if err != nil {
		log.Fatal(err)
		return constants.MESSAGE_LOGOUT_FAILED + ", Cause: " + err.Error()
	}
	data := types.Data{}
	err = json.Unmarshal(content, &data)
	if err != nil {
		log.Fatal(err)
		return constants.MESSAGE_LOGOUT_FAILED + ", Cause: " + err.Error()
	}

	customers := data.Customers
	for i, customer := range customers {
		if customer.Id == id {
			if !customer.LoggedIn {
				return constants.MESSAGE_LOGOUT_FAILED + ", Cause: " + constants.MESSAGE_NOT_LOGGED_IN_YET
			}

			customers[i].LoggedIn = false
			break
		}
	}
	data.Customers = customers

	history := "[" + utils.Now() + "] Customer with id " + fmt.Sprint(id) + " logged out"
	data.Histories = append(data.Histories, history)

	content, err = json.Marshal(data)
	if err != nil {
		log.Fatal(err)
		return constants.MESSAGE_LOGOUT_FAILED + ", Cause: " + err.Error()
	}

	err = os.WriteFile(constants.DATA_PATH, content, 0644)
	if err != nil {
		log.Fatal(err)
		return constants.MESSAGE_LOGOUT_FAILED + ", Cause: " + err.Error()
	}

	return constants.MESSAGE_LOGOUT_SUCCESS
}
