package routers

import (
	"github.com/gin-gonic/gin"
	"yumiao/routers/api"

)

//初始化路由
func InitRouter(r *gin.Engine) *gin.Engine {

	someGroup := r.Group("/api")
	{
		api.InitBooksRouter(someGroup)
		api.InitNotesRouter(someGroup)
		api.InitUsersRouter(someGroup)
		api.InitAuthRouter(someGroup)
		api.InitNoteSpacesRouter(someGroup)
	}

	return r
}

