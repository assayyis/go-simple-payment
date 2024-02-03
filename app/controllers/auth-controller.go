package controllers

import (
	"assessment/merchant-bank-payment/app/constants"
	"assessment/merchant-bank-payment/app/services"
	"assessment/merchant-bank-payment/app/types"

	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func Login(c *gin.Context) {
	var credential types.Customer

	if err := c.BindJSON(&credential); err != nil {
		return
	}

	msg := services.Login(credential.Username, credential.Password)

	if constants.MESSAGE_LOGIN_SUCCESS == msg {
		c.IndentedJSON(http.StatusOK, msg)
	} else {
		c.IndentedJSON(http.StatusInternalServerError, msg)
	}
}

func Logout(c *gin.Context) {
	id := c.Param("id")
	if n, err := strconv.Atoi(id); err == nil {
		msg := services.Logout(n)

		if constants.MESSAGE_LOGOUT_SUCCESS == msg {
			c.IndentedJSON(http.StatusOK, msg)
		} else {
			c.IndentedJSON(http.StatusInternalServerError, msg)
		}
	}
}
