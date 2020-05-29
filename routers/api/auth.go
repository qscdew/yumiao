package api

import (
	"github.com/gin-gonic/gin"
	"yumiao/controllers"
)

func InitAuthRouter(Group *gin.RouterGroup )  {

	Group.GET("/auth", controllers.GetAuth)
	Group.POST("/auth/login", controllers.Login)
	Group.POST("/auth/register", controllers.Register)
	Group.GET("/auth/logout", controllers.Logout)
}