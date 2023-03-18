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
	merchant := r.Router.Group("merchant")
	merchant.GET("/profile", middleware.ValidateToken(), r.MerchantHandler.GetMe)
}
