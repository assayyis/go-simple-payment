package app

import (
	"assessment/merchant-bank-payment/app/constants"
	"assessment/merchant-bank-payment/app/controllers"

	"github.com/gin-gonic/gin"
)

func Run() {

	router := gin.Default()
	router.POST(constants.URL_LOGIN, controllers.Login)
	router.GET(constants.URL_LOGOUT, controllers.Logout)
	router.POST(constants.URL_PAYMENT, controllers.Payment)

	router.Run(constants.APP_URL)
}
