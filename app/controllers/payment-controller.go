package controllers

import (
	"assessment/merchant-bank-payment/app/constants"
	"assessment/merchant-bank-payment/app/services"
	"assessment/merchant-bank-payment/app/types"

	"net/http"

	"github.com/gin-gonic/gin"
)

func Payment(c *gin.Context) {
	var payment types.Payment

	if err := c.BindJSON(&payment); err != nil {
		return
	}

	msg := services.Pay(payment.CustomerId, payment.MerchantId, payment.Amount)

	if constants.MESSAGE_PAYMENT_SUCCESS == msg {
		c.IndentedJSON(http.StatusOK, msg)
	} else {
		c.IndentedJSON(http.StatusInternalServerError, msg)
	}
}
