package main

import (
	"fmt"
	"log"
	"os"
	"url-shortener/cache"
	"url-shortener/database"
	"url-shortener/routes"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	// Load .env file
	if err := godotenv.Load(); err != nil {
		log.Fatal("Failed to load .env file: ", err)
	}

	// Get port from environment variable
	port := os.Getenv("API_PORT")
	if port == "" {
		port = "8080" // Fallback to default port
	}

	gin.SetMode(gin.ReleaseMode)
	database.ConnectDB()
	cache.InitRedis() // Initialize Redis
	router := gin.Default()
	router.Use(func(c *gin.Context) {
		c.Next()
		log.Printf("Request: %s %s, Status: %d", c.Request.Method, c.Request.URL, c.Writer.Status())
	})
	router.POST("/shorten", routes.ShortenURL)
	router.GET("/:code", routes.RedirectURL)
	router.GET("/health", routes.Health) // Add health endpoint

	// Run server with dynamic port
	addr := fmt.Sprintf(":%s", port)
	log.Printf("Starting server on %s", addr)
	if err := router.Run(addr); err != nil {
		log.Fatal("Failed to start server: ", err)
	}
}