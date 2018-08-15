package server

import (
	"github.com/foolin/gin-template"
	"github.com/gin-contrib/static"
	"github.com/gin-gonic/gin"
	"github.com/lmorel3/guard-go/app/controllers"
	"github.com/lmorel3/guard-go/app/middlewares"
)

func CreateRouter() *gin.Engine {
	router := gin.New()
	router.HTMLRender = gintemplate.Default()

	router.Use(gin.Logger())

	// Must come before Recovery
	router.Use(middlewares.AuthMiddleware())

	router.Use(gin.Recovery())
	router.Use(static.Serve("/assets", static.LocalFile("./assets", false)))

	mainGroup := router.Group("")
	{
		mainCtrl := new(controllers.MainController)
		userCtrl := new(controllers.UsersController)
		authCtrl := new(controllers.AuthController)

		mainGroup.GET("/", mainCtrl.Index)
		mainGroup.GET("/check", authCtrl.Check)

		mainGroup.GET("/login", userCtrl.Login)
		mainGroup.POST("/login", userCtrl.LoginAttempt)

		mainGroup.GET("/logout", userCtrl.Logout)

		mainGroup.GET("/password", userCtrl.SetPassword)
		mainGroup.POST("/password", userCtrl.SetPassword)
	}

	adminGroup := router.Group("admin")
	{
		adminCtrl := new(controllers.AdminController)

		adminGroup.GET("/", adminCtrl.Index)

		userGroup := adminGroup.Group("users")
		{
			userGroup.GET("/add", adminCtrl.UserAdd)
			userGroup.POST("/add", adminCtrl.UserAdd)

			userGroup.GET("/delete/:username", adminCtrl.UserDelete)
		}

	}

	return router

}
