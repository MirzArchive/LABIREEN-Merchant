package routes

import (
	"labireen-merchant/handlers"
	"labireen-merchant/middleware"

	"github.com/gin-gonic/gin"
)

type CustomerRoutes struct {
	Router          *gin.Engine
	CustomerHandler handlers.CustomerHandler
}

func (r *CustomerRoutes) Register() {
	merchant := r.Router.Group("merchant")
	merchant.GET("/profile", middleware.ValidateToken(), r.CustomerHandler.GetMe)
}
