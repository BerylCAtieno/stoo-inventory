package main

import (
	"log"

	"github.com/gin-gonic/gin"

	"github.com/berylCAtieno/stoo-inventory/cmd/migration"
	"github.com/berylCAtieno/stoo-inventory/internal/config"
	"github.com/berylCAtieno/stoo-inventory/internal/database"
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
	}

	addr := ":" + config.Config.Port
	log.Println("Server running on", addr)
	err := router.Run(addr)
	if err != nil {
		log.Fatal("Server failed to start:", err)
	}
}
