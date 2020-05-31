package middlewares

import (
	"github.com/gin-gonic/gin"
	"strconv"
	"yumiao/app/jwt"
	"yumiao/models"
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

//用户只能访问属于自己的数据
func UserSelfRequired() gin.HandlerFunc {
	return func(context *gin.Context) {
		id, _ := strconv.Atoi(context.Param("id"))//用户id
		token := context.Query("token")
		claims, _ := jwt.ParseToken(token)
		if models.GetUser(id).Username==claims.Username{
			context.Next()
		}else{
			context.Abort()

			context.JSON(402, gin.H{
				"message": "No access",
			})
		}
}
}

//用户只能访问属于自己的数据
func UserSelfNoteSpace() gin.HandlerFunc {
	return func(context *gin.Context) {
		id, _ := strconv.Atoi(context.Param("id"))//notespace id
		token := context.Query("token")
		claims, _ := jwt.ParseToken(token)
		if models.GetUser(int(models.GetNoteSpace(id).UserID)).Username==claims.Username{
			context.Next()
		}else{
			context.Abort()

			context.JSON(402, gin.H{
				"message": "No access",
			})
		}
	}
}

//用户只能访问属于自己的数据
func UserSelfNote() gin.HandlerFunc {
	return func(context *gin.Context) {
		id, _ := strconv.Atoi(context.Param("id"))//noteid
		token := context.Query("token")
		claims, _ := jwt.ParseToken(token)

		if models.GetUser(int(models.GetNote(id).UserID)).Username==claims.Username{
			context.Next()
		}else{
			context.Abort()

			context.JSON(402, gin.H{
				"message": "No access",
			})
		}
	}
}