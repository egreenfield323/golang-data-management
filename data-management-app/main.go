package main

import (
	"log"

	"data-management-app/db"
	"data-management-app/handlers"

	// use gin web framework
	"github.com/gin-gonic/gin"
)

func main() {
	if err := db.InitDB(); err != nil {
		log.Fatalf("Failed to initialize database: %v", err)
	}

	r := gin.Default()

	// Handles CORS errors in html console
	r.Use(func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type")
		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}
		c.Next()
	})

	handlers.RegisterUserRoutes(r)

	// Start server on port 8080
	if err := r.Run(":8080"); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
