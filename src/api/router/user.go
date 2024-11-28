package router

import (
	"github.com/garickedd/ReLibca/src/api/config"
	"github.com/garickedd/ReLibca/src/api/handler"
	"github.com/garickedd/ReLibca/src/api/middleware"
	"github.com/gin-gonic/gin"
)

func User(router *gin.RouterGroup, cfg *config.Config) {
	h := handler.NewUserHandler(cfg)

	router.POST("/send-otp", middleware.OtpLimiter(cfg), h.SendOtp)
	router.POST("/login-by-username", h.LoginByUsername)
	router.POST("/register-by-username", h.RegisterByUsername)
	router.POST("/login-by-mobile", h.RegisterLoginByMobileNumber)
}
