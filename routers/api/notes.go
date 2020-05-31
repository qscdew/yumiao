package api

import (
	"github.com/gin-gonic/gin"
	"yumiao/controllers"
	"yumiao/middlewares"
)

func InitNotesRouter(Group *gin.RouterGroup )  {
	Group.GET("/notes", controllers.GetNotes)
	Group.GET("/notes/:id", middlewares.UserSelfNote(),controllers.GetNote)

	Group.DELETE("/notes/:id", middlewares.UserSelfNote(),controllers.DeleteNote)
	Group.PUT("/notes/:id", middlewares.UserSelfNote(),controllers.EditNote)

}

