package handlers

import (
	"context"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v8"
	"github.com/sourabhsd87/URL_Shortner/config"
	"github.com/sourabhsd87/URL_Shortner/db"
	"github.com/sourabhsd87/URL_Shortner/models"
)

func Redirect(c *gin.Context) {
	config.Logger.Debug("")
	shortCode := c.Param("shortCode")

	longURL, err := config.RedisClient.Get(context.Background(), shortCode).Result()
	if err == redis.Nil {
		var url models.URL
		if err := db.DB.Where("short_url = ?", shortCode).First(&url).Error; err != nil {
			c.JSON(http.StatusNotFound, models.Response{
				StatusCode: http.StatusNotFound,
				Message:    "URL not found",
				Data:       map[string]interface{}{},
			})
			return
		}
		config.RedisClient.Set(context.Background(), shortCode, url.LongURL, 24*time.Hour)
		longURL = url.LongURL

	}
	c.Redirect(http.StatusFound, longURL)
}
