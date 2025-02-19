package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/sourabhsd87/URL_Shortner/handlers"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()
	r.POST("/shorten", handlers.ShortenURL)
	r.GET("/:shortCode", handlers.Redirect)
	return r
}
