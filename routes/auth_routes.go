package routes

import (
	"labireen-merchant/handlers"

	"github.com/gin-gonic/gin"
)

type AuthRoutes struct {
	Router      *gin.Engine
	AuthHandler handlers.AuthHandler
}

func (r *AuthRoutes) Register() {
	auth := r.Router.Group("auth")
	auth.POST("/register", r.AuthHandler.RegisterMerchant)
	auth.POST("/login", r.AuthHandler.LoginMerchant)
	auth.GET("/verify/:verification-code", r.AuthHandler.VerifyEmail)
}
