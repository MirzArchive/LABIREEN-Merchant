package routes

import (
	"labireen-merchant/handlers"
	"labireen-merchant/middleware"

	"github.com/gin-gonic/gin"
)

type MerchantRoutes struct {
	Router          *gin.Engine
	MerchantHandler handlers.MerchantHandler
}

func (r *MerchantRoutes) Register() {
	Merchant := r.Router.Group("Merchant")
	Merchant.GET("/profile", middleware.ValidateToken(), r.MerchantHandler.GetMe)
}
