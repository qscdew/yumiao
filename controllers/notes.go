package controllers

import (

	"github.com/gin-gonic/gin"
	"strconv"
	"yumiao/models"
)

//获取全部笔记
func GetNotes(c *gin.Context) {
	data:=models.GetAllNotes()
	c.JSON(200, gin.H{
		"message": "success",
		"data":data,
	})
}

//获取指定笔记
func GetNote(c *gin.Context) {
	id ,_:= strconv.Atoi(c.Param("id"))
	data:=models.GetNote(id)
	c.JSON(200, gin.H{
		"message": data,
	})

}

//删除指定笔记
func DeleteNote(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	models.DeleteNote(id)
	c.JSON(200, gin.H{
		"message": "success",
	})
}

//修改指定笔记
func EditNote(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	var form models.Note
	if c.Bind(&form)==nil{
		if models.EditNote(id,form){
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


