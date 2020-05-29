package middlewares

import (
	"github.com/gin-gonic/gin"
)

func LoginRequired() gin.HandlerFunc {
	return func(context *gin.Context) {
		cookie, err := context.Request.Cookie("user_cookie")
		if err == nil {
			context.SetCookie(cookie.Name, cookie.Value, 1000, cookie.Path, cookie.Domain, cookie.Secure, cookie.HttpOnly)

			context.Next()
		} else {
			context.Abort()
			context.JSON(401, gin.H{
				"message": "login required",
			})
		}

	}
}

