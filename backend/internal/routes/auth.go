package routes

import (
	"github.com/berylCAtieno/stoo-inventory/internal/handlers"

	"github.com/berylCAtieno/stoo-inventory/internal/middleware"

	"github.com/gin-gonic/gin"
)

func AuthRoutes(r *gin.Engine, authHandler *handlers.AuthHandler) {
	auth := r.Group("/auth")
	{
		auth.POST("/register", authHandler.Register)
		auth.POST("/verify-otp", authHandler.VerifyOTP)
		auth.POST("/login", authHandler.Login)
		// Forgot/Reset password
		auth.POST("/forgot-password", authHandler.ForgotPassword)
		auth.POST("/reset-password", authHandler.ResetPassword)
	}

	// Protected example
	protected := r.Group("/protected")
	protected.Use(middleware.AuthRequired())
	{
		protected.GET("/profile", func(c *gin.Context) {
			userID, _ := c.Get("userID")
			c.JSON(200, gin.H{"message": "Protected route", "userID": userID})
		})
	}
}
