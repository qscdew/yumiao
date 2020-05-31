package routers

import (
	"github.com/gin-gonic/gin"
	"yumiao/controllers"
	"yumiao/middlewares/jwt"
	"yumiao/routers/api"
)

//初始化路由
func InitRouter(r *gin.Engine) *gin.Engine {

	r.GET("/api/auth", controllers.GetAuth)
	someGroup := r.Group("/api")
	someGroup.Use(jwt.JWT())
	{
		api.InitBooksRouter(someGroup)
		api.InitNotesRouter(someGroup)
		api.InitUsersRouter(someGroup)
		api.InitNoteSpacesRouter(someGroup)
	}

	return r
}

