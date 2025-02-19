package handlers

import (
	"context"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/sourabhsd87/URL_Shortner/config"
	"github.com/sourabhsd87/URL_Shortner/db"
	"github.com/sourabhsd87/URL_Shortner/models"
)

type request struct {
	LongURL string `json:"long_url"`
}

func ShortenURL(c *gin.Context) {
	config.Logger.Debug("In ShortenURL handler")

	var req request
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, models.Response{
			StatusCode: http.StatusBadRequest,
			Message:    "Invalid request",
			Data:       map[string]interface{}{},
		})
		config.Logger.Error(err.Error())
		return
	}

	shortCode := uuid.New().String()[:6]

	url := models.URL{
		LongURL:  req.LongURL,
		ShortURL: shortCode,
		Expiry:   time.Now().Add(34 * time.Hour),
	}

	if err := db.DB.Create(&url).Error; err != nil {
		c.JSON(http.StatusInternalServerError, models.Response{
			StatusCode: http.StatusInternalServerError,
			Message:    "Error creating URL",
			Data:       map[string]interface{}{},
		})
		config.Logger.Error(err.Error())
		return
	}
	config.RedisClient.Set(context.Background(), shortCode, url.LongURL, 24*time.Hour)

	c.JSON(http.StatusOK, models.Response{
		StatusCode: http.StatusOK,
		Message:    "URL shortened successfully",
		Data: map[string]interface{}{
			"short_url": "http://" + config.Host + ":" + config.Port + "/" + shortCode,
		},
	})
}
