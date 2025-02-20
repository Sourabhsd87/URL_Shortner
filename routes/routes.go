package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/sourabhsd87/URL_Shortner/handlers"
	"github.com/sourabhsd87/URL_Shortner/middlewares"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()

	r.GET("/auth/google", handlers.HandleGoogleLogin)
	r.GET("/auth/callback", handlers.HandleGoogleCallback)

	OAuth := r.Group("/auth", middlewares.AuthMiddleware())
	OAuth.POST("/shorten", handlers.ShortenURL)
	r.GET("/:shortCode", handlers.Redirect)
	return r

}
