package router

import (
	"github.com/Fillybodyknow/blog-api/internal/handler"
	"github.com/Fillybodyknow/blog-api/internal/middleware"
	"github.com/gin-gonic/gin"
)

type AuthRouter struct {
	AuthHandler *handler.AuthHandler
}

func NewAuthRouter(authHandler *handler.AuthHandler) *AuthRouter {
	return &AuthRouter{AuthHandler: authHandler}
}

func (r *AuthRouter) AuthRoutes(rg *gin.RouterGroup) {
	rg.POST("/register", r.AuthHandler.RegisterUser)
	rg.POST("/login", r.AuthHandler.LoginUser)
}

func (r *AuthRouter) OTPRoutes(rg *gin.RouterGroup) {
	rg.Use(middleware.AuthMiddleware())
	rg.GET("/verify_otp", r.AuthHandler.OTP)
	rg.POST("/verify_otp", r.AuthHandler.VerifyOTP)
}
