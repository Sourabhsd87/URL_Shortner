package config

import (
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
)

var GoogleOAuthConfig *oauth2.Config

func InitOAuth() {
	GoogleOAuthConfig = &oauth2.Config{
		ClientID:     Google_Client_ID,
		ClientSecret: Google_Client_Secret,
		RedirectURL:  Google_Redirect_URL,
		Scopes:       []string{"https://www.googleapis.com/auth/userinfo.email", "https://www.googleapis.com/auth/userinfo.profile"},
		Endpoint:     google.Endpoint,
	}
}
