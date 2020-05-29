package api
import (
	"github.com/gin-gonic/gin"
	"yumiao/controllers"
)

func InitNoteSpacesRouter(Group *gin.RouterGroup )  {
	Group.GET("/notespaces", controllers.GetNoteSpaces)
	Group.GET("/notespaces/:id", controllers.GetNoteSpace)
	Group.DELETE("/notespaces/:id", controllers.DeleteNoteSpace)
	Group.PUT("/notespaces/:id", controllers.EditNoteSpace)

	Group.GET("/notespaces/:id/notes", controllers.GetNoteSpaceNotes)
	Group.POST("/notespaces/:id/notes", controllers.AddNote)

}


