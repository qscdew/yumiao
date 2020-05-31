package api

import (
	"github.com/gin-gonic/gin"
	"yumiao/controllers"
	"yumiao/middlewares"
)

func InitUsersRouter(Group *gin.RouterGroup )  {
	Group.GET("/users", controllers.GetUsers)
	Group.GET("/users/:id", controllers.GetUser)
	Group.GET("/users/:id/notespaces", middlewares.UserSelfRequired(),controllers.GetUserNoteSpaces)
	Group.POST("/users/:id/notespaces",  middlewares.UserSelfRequired(),controllers.UserAddNoteSpace)

}


