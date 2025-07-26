package routes

import (
	"log"
	"fmt"
	"time"
	"net/http"
	"url-shortener/database"
	"url-shortener/models"
	"url-shortener/cache"
	"github.com/redis/go-redis/v9"
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
	start := time.Now()
	
	// Check Redis cache
    originalURL, err := cache.GetURL(code)
	log.Printf("Redis query for %s took %v", code, time.Since(start))

    if err == nil {
        // Cache hit
        c.Redirect(http.StatusFound, originalURL)
        return
    }
    if err != redis.Nil {
        // Log Redis errors but proceed to database
        log.Printf("Redis error for short_code %s: %v", code, err)
    }

	// Cache miss, query the database
	var url models.URL
	start = time.Now()
	if err := database.DB.Where("short_code = ?", code).Take(&url).Error; err != nil {
		log.Printf("Database error for short_code %s: %v, took %v", code, err, time.Since(start))
		c.JSON(http.StatusNotFound, gin.H{"error": "URL not found"})
		return
	}

	// Cache the result
    if err := cache.SetURL(code, url.OriginalURL, 1*time.Hour); err != nil {
        log.Printf("Failed to cache short_code %s: %v", code, err)
    }

	c.Redirect(http.StatusFound, url.OriginalURL)
}

// Health handles GET /health to return a status check
func Health(c *gin.Context) {
    c.JSON(http.StatusOK, gin.H{"status": "ok"})
}
