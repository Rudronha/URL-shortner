package routes

import (
	"log"
	"fmt"
	"net/http"
	"url-shortener/database"
	"url-shortener/models"

	"github.com/gin-gonic/gin"
	"github.com/teris-io/shortid"
)

func ShortenURL(c *gin.Context) {
	var req struct {
		URL string `json:"url"`
	}

	if err := c.ShouldBindJSON(&req); err != nil || req.URL == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid URL"})
		return
	}

	code, _ := shortid.Generate()

	newURL := models.URL{
		ShortCode:   code,
		OriginalURL: req.URL,
	}

	if err := database.DB.Create(&newURL).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save URL"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"short_url": fmt.Sprintf("http://localhost:8080/%s", code)})
}

func RedirectURL(c *gin.Context) {
	code := c.Param("code")
	var url models.URL

	if err := database.DB.Where("short_code = ?", code).Take(&url).Error; err != nil {
		log.Printf("Database error for short_code %s: %v", code, err)
		c.JSON(http.StatusNotFound, gin.H{"error": "URL not found"})
		return
	}

	c.Redirect(http.StatusFound, url.OriginalURL)
}
