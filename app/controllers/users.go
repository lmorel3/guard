package controllers

import (
	"log"
	"net/http"
	"net/url"

	"github.com/gin-gonic/gin"
	"github.com/xyproto/permissionbolt"

	"github.com/lmorel3/guard-go/app/config"
)

type UsersController struct{}

func (u UsersController) Login(c *gin.Context) {
	userState := c.MustGet("userstate").(*permissionbolt.UserState)
	cookieUser := userState.Username(c.Request)

	if userState.IsLoggedIn(cookieUser) {
		log.Println("Already logged in")
		c.Redirect(http.StatusSeeOther, "/")
		return
	}

	c.HTML(http.StatusOK, "login", gin.H{"badCredentials": false})
}

func (u UsersController) LoginAttempt(c *gin.Context) {
	userState := c.MustGet("userstate").(*permissionbolt.UserState)

	success := true

	username := c.PostForm("username")
	password := c.PostForm("password")

	if len(username) > 0 && len(password) > 0 {
		config := config.GetConfig()
		guardURL := "http://" + config.GetString("guard")

		redirectURL, err := url.QueryUnescape(c.Query("url"))
		if nil != err {
			redirectURL = guardURL
		}

		// Checks if username exists
		success = userState.CorrectPassword(username, password)
		if success {
			userState.Login(c.Writer, username)

			// Set the right "domain" for the cookie
			SetCookieDomain(c)

			log.Printf("Redirecting to " + redirectURL)
			c.Redirect(http.StatusSeeOther, redirectURL)
			return
		}
	}

	c.HTML(http.StatusOK, "login", gin.H{"badCredentials": !success})
}

func (u UsersController) Logout(c *gin.Context) {
	userState := c.MustGet("userstate").(*permissionbolt.UserState)
	cookieUser := userState.Username(c.Request)

	if len(cookieUser) > 0 {
		log.Println("User is logged in", cookieUser)
		userState.Logout(cookieUser)
		userState.ClearCookie(c.Writer)

		SetCookieDomain(c)
	}

	c.Redirect(http.StatusTemporaryRedirect, "/")
}

func (u UsersController) SetPassword(c *gin.Context) {
	userState := c.MustGet("userstate").(*permissionbolt.UserState)

	username := userState.Username(c.Request)
	oldPassword := c.PostForm("old")
	newPassword := c.PostForm("new")

	success := false

	if len(oldPassword) > 0 && len(newPassword) > 0 {
		if userState.CorrectPassword(username, oldPassword) {
			userState.SetPassword(username, newPassword)
			success = true
		}
	}

	c.HTML(http.StatusOK, "password", gin.H{"updated": success})
}

func SetCookieDomain(c *gin.Context) {
	cookie := c.Writer.Header().Get("Set-Cookie")
	cookie = cookie + "; Domain=" + config.GetConfig().GetString("domain")

	c.Writer.Header().Set("Set-Cookie", cookie)
}
