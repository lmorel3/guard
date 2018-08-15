package middlewares

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/xyproto/permissionbolt"
)

func AuthMiddleware() gin.HandlerFunc {

	// New permissions middleware
	perm, err := permissionbolt.NewWithConf("../config/bolt.db")
	if err != nil {
		log.Fatalln(err)
	}

	//	perm.AddUserPath("/check")
	perm.AddUserPath("/password")

	perm.AddPublicPath("/assets")
	perm.AddPublicPath("/check")

	userState := perm.UserState()

	usernames, err2 := userState.AllUsernames()
	if err2 != nil || len(usernames) == 0 {
		log.Println("Creating default 'admin/admin' user")
		userState.AddUser("admin", "admin", "admin@guard.local")
		userState.SetAdminStatus("admin")
	}

	return func(c *gin.Context) {
		cookieUser := userState.Username(c.Request)
		isLogged := userState.IsLoggedIn(cookieUser)

		// Happens when app has been shutdown and user still has a cookie
		if len(cookieUser) > 0 && !isLogged {
			log.Printf("% should be logged in since its cookie is valid: marked as logged in", cookieUser)
			userState.SetLoggedIn(cookieUser)
		}

		// Check if the user has the right admin/user rights
		if perm.Rejected(c.Writer, c.Request) {
			// Deny the request, don't call other middleware handlers
			c.AbortWithStatus(http.StatusForbidden)
			fmt.Fprint(c.Writer, "Permission denied!")
			return
		}

		// Call the next middleware handler
		c.Set("userstate", userState)
		c.Next()

	}
}
