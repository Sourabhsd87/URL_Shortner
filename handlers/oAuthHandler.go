package handlers

import (
	"context"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sourabhsd87/URL_Shortner/config"
	"golang.org/x/oauth2"
)

// OAuth login handler
func HandleGoogleLogin(c *gin.Context) {
	url := config.GoogleOAuthConfig.AuthCodeURL("state-token", oauth2.AccessTypeOffline)
	c.Redirect(http.StatusFound, url)
}

// OAuth callback handler
func HandleGoogleCallback(c *gin.Context) {
	code := c.Query("code")
	if code == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Authorization code not provided"})
		return
	}

	// Exchange code for token
	token, err := config.GoogleOAuthConfig.Exchange(context.Background(), code)
	if err != nil {
		fmt.Println("")
		config.Logger.Error("Error exchanging code for token: " + err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Token exchange failed"})
		return
	}

	// Get user info from Google
	client := config.GoogleOAuthConfig.Client(context.Background(), token)
	resp, err := client.Get("https://www.googleapis.com/oauth2/v2/userinfo")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get user info"})
		return
	}
	defer resp.Body.Close()

	// Store user session in Redis
	sessionKey := "session:" + token.AccessToken
	config.RedisClient.Set(config.Ctx, sessionKey, token.AccessToken, 0) // Store session indefinitely

	c.JSON(http.StatusOK, gin.H{"message": "Login successful", "token": token.AccessToken})
}
