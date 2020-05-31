package controllers

import (
	"github.com/gin-gonic/gin"
	"strconv"
	"yumiao/models"
)

//获取全部笔记
func GetNoteSpaces(c *gin.Context) {
	var data interface {}
	data=models.GetAllNoteSpaces()
	c.JSON(200, gin.H{
		"message": "success",
		"data":data,
	})
}

//获取指定笔记区域
func GetNoteSpace(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	data:=models.GetNoteSpace(id)


	c.JSON(200, gin.H{
		"message":"success",
		"data":data,
	})
}

//修改指定笔记区域
func EditNoteSpace(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	var form models.NoteSpace
	if c.Bind(&form)==nil{
		if models.EditNoteSpace(id,form){
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
//删除指定笔记区域
func DeleteNoteSpace(c *gin.Context) {

	id, _ := strconv.Atoi(c.Param("id"))
	models.DeleteNoteSpace(id)

	c.JSON(200, gin.H{
		"message":"success",

	})
}

//获取指定笔记区域的所有笔记
func GetNoteSpaceNotes(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))


	c.JSON(200, gin.H{
		"message":"success",
		"data":models.GetNoteSpaceNotes(id),
	})

}

//为某指定笔记区域添加笔记
func AddNote(c *gin.Context){
	var form models.Note
	id, _ := strconv.Atoi(c.Param("id"))
	form.NoteSpaceID= uint(id)
	if c.Bind(&form)==nil{
		if models.AddNote(form){
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
