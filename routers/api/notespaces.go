package api
import (
	"github.com/gin-gonic/gin"
	"yumiao/controllers"
	"yumiao/middlewares"
)

func InitNoteSpacesRouter(Group *gin.RouterGroup )  {
	Group.GET("/notespaces", controllers.GetNoteSpaces)
	Group.GET("/notespaces/:id", middlewares.UserSelfNoteSpace(),controllers.GetNoteSpace)
	Group.DELETE("/notespaces/:id",middlewares.UserSelfNoteSpace(), controllers.DeleteNoteSpace)
	Group.PUT("/notespaces/:id", middlewares.UserSelfNoteSpace(),controllers.EditNoteSpace)

	Group.GET("/notespaces/:id/notes", middlewares.UserSelfNoteSpace(),controllers.GetNoteSpaceNotes)
	Group.POST("/notespaces/:id/notes",middlewares.UserSelfNoteSpace(), controllers.AddNote)

}


