package handlers

import (
	"context"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v8"
	"github.com/sirupsen/logrus"
	"github.com/sourabhsd87/URL_Shortner/config"
	"github.com/sourabhsd87/URL_Shortner/db"
	"github.com/sourabhsd87/URL_Shortner/models"
)

func Redirect(c *gin.Context) {
	config.Logger.Debug("Redirect function called")
	shortCode := c.Param("shortCode")
	config.Logger.Debug("Short code: " + shortCode)

	longURL, err := config.RedisClient.Get(context.Background(), shortCode).Result()
	if err == redis.Nil {
		config.Logger.Warn("Short code not found in Redis")
		var url models.URL
		if err := db.DB.Where("short_url = ?", shortCode).First(&url).Error; err != nil {
			config.Logger.Error("URL not found in database: " + err.Error())
			c.JSON(http.StatusNotFound, models.Response{
				StatusCode: http.StatusNotFound,
				Message:    "URL not found",
				Data:       map[string]interface{}{},
			})
			return
		}
		config.RedisClient.Set(context.Background(), shortCode, url.LongURL, 24*time.Hour)
		longURL = url.LongURL

		url.Clicks++
		config.Logger.WithFields(logrus.Fields{
			"URL":    url.LongURL,
			"clicks": url.Clicks,
		})

		if err := db.DB.Save(&url).Error; err != nil {
			config.Logger.Error("Error updating URL clicks: " + err.Error())
		}

		config.Logger.Info("URL fetched from database and cached in Redis")
	}
	config.Logger.Info("Redirecting to long URL: " + longURL)
	c.Redirect(http.StatusFound, longURL)
}
