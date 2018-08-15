package controllers

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/xyproto/permissionbolt"
)

type AdminController struct{}

func (u AdminController) Index(c *gin.Context) {
	userState := c.MustGet("userstate").(*permissionbolt.UserState)
	usernames, err := userState.AllUsernames()

	if nil != err {
		usernames = []string{}
	}

	c.HTML(http.StatusOK, "admin_users", gin.H{
		"usernames": usernames,
		"isLoggedIn": func(username string) bool {
			return userState.IsLoggedIn(username)
		},
		"isAdmin": func(username string) bool {
			return userState.IsAdmin(username)
		},
	})
}

func (u AdminController) UserAdd(c *gin.Context) {
	userState := c.MustGet("userstate").(*permissionbolt.UserState)

	username := c.PostForm("username")
	password := c.PostForm("password")
	email := username + "@local"

	makeAdmin := len(c.PostForm("make_admin")) > 0

	alreadyTaken := false
	if userState.HasUser(username) {
		alreadyTaken = true
	} else if len(username) > 0 {
		userState.AddUser(username, password, email)
		userState.Confirm(username)

		if makeAdmin {
			userState.SetAdminStatus(username)
		}

		c.Redirect(http.StatusSeeOther, "/admin")
		return
	}

	c.HTML(http.StatusOK, "admin_user_add", gin.H{
		"alreadyTaken": alreadyTaken,
	})
}

func (u AdminController) UserDelete(c *gin.Context) {
	userState := c.MustGet("userstate").(*permissionbolt.UserState)
	cookieUser := userState.Username(c.Request)

	username := c.Param("username")

	if cookieUser != username {
		userState.RemoveUser(username)
	} else {
		log.Printf("%s tries to delete himself...", username)
	}

	c.Redirect(http.StatusSeeOther, "/admin")
}
