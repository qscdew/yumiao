package api

import (
	"github.com/gin-gonic/gin"
	"yumiao/controllers"
)

func InitNotesRouter(Group *gin.RouterGroup )  {
	Group.GET("/notes", controllers.GetNotes)
	Group.GET("/notes/:id", controllers.GetNote)

	Group.DELETE("/notes/:id", controllers.DeleteNote)
	Group.PUT("/notes/:id", controllers.EditNote)

}

