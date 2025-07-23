package main

import (
	"url-shortener/database"
	"url-shortener/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	gin.SetMode(gin.ReleaseMode)
	database.ConnectDB()

	router := gin.Default()
	router.POST("/shorten", routes.ShortenURL)
	router.GET("/:code", routes.RedirectURL)

	router.Run(":8080")
}
