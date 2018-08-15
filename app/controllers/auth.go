package controllers

import (
	"log"
	"net/http"
	"net/url"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"github.com/xyproto/permissionbolt"

	"github.com/lmorel3/guard-go/app/config"
)

type AuthController struct{}

func (u AuthController) Check(c *gin.Context) {
	h := c.Request.Header
	requestURL := h.Get("X-Forwarded-Host") + h.Get("X-Forwarded-Uri")
	config := config.GetConfig()

	if isAllowed(requestURL, config, c) {

		// User is already logged in
		log.Println("Logged in or auth path: ok")
		c.AbortWithStatus(http.StatusNoContent)

	} else if isPubliclyAllowed(requestURL, config) {

		log.Println("Publicly allowed: ok")
		// User is not logged in but resource is publicly available
		c.AbortWithStatus(http.StatusNoContent)

	} else {

		// User must login
		guardURL := config.GetString("guard")
		requestURL = url.QueryEscape(h.Get("X-Forwarded-Proto") + "://" + requestURL)
		redirectURL := "http://" + guardURL + "/login?url=" + requestURL

		log.Println("User must login " + redirectURL)

		c.Redirect(http.StatusTemporaryRedirect, redirectURL)
	}
}

func isPubliclyAllowed(requestURL string, config *viper.Viper) bool {
	allowedUrls := config.GetStringSlice("allowed")

	for _, publicURL := range allowedUrls {
		if strings.HasPrefix(requestURL, publicURL) {
			return true
		}
	}

	return false
}

func isAllowed(requestURL string, config *viper.Viper, c *gin.Context) bool {
	userState := c.MustGet("userstate").(*permissionbolt.UserState)
	cookieUser := userState.Username(c.Request)

	guardURL := config.GetString("guard")

	return userState.IsLoggedIn(cookieUser) || strings.HasPrefix(requestURL, guardURL)
}
