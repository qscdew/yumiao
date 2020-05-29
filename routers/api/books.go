package api

import (
	"github.com/gin-gonic/gin"
	"yumiao/controllers"
	"yumiao/middlewares"
)

func InitBooksRouter(Group *gin.RouterGroup )  {
	Group.GET("/books",controllers.GetBooks)
	Group.GET("/books/:id", controllers.GetBook)
	Group.POST("/books", middlewares.LoginRequired(),controllers.AddBook)
	Group.DELETE("/books/:id", middlewares.LoginRequired(),controllers.DeleteBook)
	Group.PUT("/books/:id", middlewares.LoginRequired(),controllers.EditBook)

}