package routes

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/inontz/kepler-api/handlers"
	"gopkg.in/resty.v1"
)

// AuthSuccess get access token
type AuthSuccess struct {
	AccessToken  string  `json:"access_token"`
	ExpiresIn    float64 `json:"expires_in"`
	IDToken      string  `json:"id_token"`
	RefreshToken string  `json:"refresh_token"`
	Scope        string  `json:"scope"`
	TokenType    string  `json:"token_type"`
}

// Profile Gets a user's display name, profile image, and status message.
// See: https://developers.line.biz/en/reference/social-api/#get-user-profile
type Profile struct {
	DisplayName   string `json:"displayName"`
	UserID        string `json:"userId"`
	PictureURL    string `json:"pictureUrl"`
	StatusMessage string `json:"statusMessage"`
}

func RegisterRouter(r *gin.RouterGroup) {
	r.Static("/static", "static")
	r.StaticFile("favicon.ico", "./resources/favicon.ico")
	r.GET("/", handlers.Index)
	r.GET("/callback", func(c *gin.Context) {
		code := c.Query("code")
		// state := c.Query("state")
		// changed := c.Query("friendship_status_changed")

		authSuccess := &AuthSuccess{}
		resp, _ := resty.R().
			SetHeader("Content-Type", "application/x-www-form-urlencoded").
			SetFormData(map[string]string{
				"grant_type":    "authorization_code",
				"code":          code,
				"redirect_uri":  "https://kepler.inontz.me",
				"client_id":     "1660948571",
				"client_secret": "19dcd1739168b460fb171c63146c5f98",
			}).
			SetResult(authSuccess). // or SetResult(AuthSuccess{}).
			Post("https://api.line.me/oauth2/v2.1/token")
		fmt.Printf("\nResponse Body: %v", resp)

		profile := &Profile{}
		resp, _ = resty.R().
			SetHeader("Authorization", "Bearer "+authSuccess.AccessToken).
			SetResult(profile). // or SetResult(AuthSuccess{}).
			Get("https://api.line.me/v2/profile")
		fmt.Printf("\nResponse Body: %v", resp)

		c.HTML(http.StatusOK, "success.html", gin.H{
			"title":         "Line QR Code Login Example",
			"userID":        profile.UserID,
			"displayName":   profile.DisplayName,
			"pictureURL":    profile.PictureURL,
			"statusMessage": profile.StatusMessage,
		})
		
	})
	// for nostr NIP-05
	// r.GET("/.well-known/nostr.json", handler.Cors, handler.NIP05)
}
