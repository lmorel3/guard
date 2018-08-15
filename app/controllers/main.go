package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/xyproto/permissionbolt"
)

type MainController struct{}

//var userModel = new(models.User)

func (u MainController) Index(c *gin.Context) {
	userState := c.MustGet("userstate").(*permissionbolt.UserState)
	username := userState.Username(c.Request)

	c.HTML(http.StatusOK, "index", gin.H{
		"username": username,
		"isLogged": (len(username) > 0),
		"isAdmin":  userState.IsAdmin(username),
	})
}
