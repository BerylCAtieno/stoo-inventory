package main

import (
	"log"

	"github.com/gin-gonic/gin"

	"github.com/berylCAtieno/stoo-inventory/cmd/migration"
	"github.com/berylCAtieno/stoo-inventory/internal/config"
	"github.com/berylCAtieno/stoo-inventory/internal/database"
	"github.com/berylCAtieno/stoo-inventory/internal/handlers"
	"github.com/berylCAtieno/stoo-inventory/internal/repositories"
	"github.com/berylCAtieno/stoo-inventory/internal/routes"
	"github.com/berylCAtieno/stoo-inventory/internal/services"
)

func main() {
	config.LoadConfig()
	database.Connect()
	migration.Run()

	router := gin.Default()

	router.GET("/", func(c *gin.Context) {
		c.String(200, "Welcome to Stoo Inventory Management App")
	})

	api := router.Group("/api/v1")
	{
		api.GET("/health", func(c *gin.Context) {
			c.JSON(200, gin.H{"status": "OK"})
		})

		userRepo := repositories.NewUserRepository(database.DB)
		authService := services.NewAuthService(userRepo)
		authHandler := handlers.NewAuthHandler(authService)

		routes.RegisterAuthRoutes(api, authHandler)
	}

	addr := ":" + config.Config.Port
	log.Println("Server running on", addr)
	if err := router.Run(addr); err != nil {
		log.Fatal("Server failed to start:", err)
	}
}
