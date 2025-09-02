package routes

import (
	"github.com/berylCAtieno/stoo-inventory/internal/handlers"

	"github.com/gin-gonic/gin"
)

func RegisterAuthRoutes(r *gin.RouterGroup, authHandler *handlers.AuthHandler) {
	auth := r.Group("/auth")
	{
		auth.POST("/register", authHandler.Register)
		auth.POST("/verify-otp", authHandler.VerifyOTP)
		auth.POST("/login", authHandler.Login)
		auth.POST("/forgot-password", authHandler.ForgotPassword)
		auth.POST("/reset-password", authHandler.ResetPassword)
	}

}
