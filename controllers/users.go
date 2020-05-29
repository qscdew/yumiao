package controllers

import (

	"github.com/gin-gonic/gin"
	"strconv"
	"yumiao/models"
)

//获取全部用户
func GetUsers(c *gin.Context) {

	c.JSON(200, gin.H{
		"message": "success",
		"data":models.GetAllUser(),
	})
}

//获取指定用户
func GetUser(c *gin.Context) {
	id ,_:= strconv.Atoi(c.Param("id"))

	c.JSON(200, gin.H{
		"message": "success",
		"data":models.GetUser(id),
	})

}

//获取指定用户的笔记区
func GetUserNoteSpaces(c *gin.Context) {
	id ,_:= strconv.Atoi(c.Param("id"))

	c.JSON(200, gin.H{
		"message": "success",
		"user":models.GetUser(id),
		"data":models.GetUserNotespaces(id),
	})

}

func UserAddNoteSpace(c *gin.Context){
	id ,_:= strconv.Atoi(c.Param("id"))

	var form models.NoteSpace
	if c.Bind(&form)==nil{
		form.UserID= uint(id)
		if models.UserAddNoteSpace(form){
			c.JSON(200, gin.H{
				"message": "success",

			})
		}else{
			c.JSON(400, gin.H{
				"message": "failed",
			})
		}
	}

}


