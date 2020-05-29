package controllers

import (
	"github.com/gin-gonic/gin"
	"strconv"
	"yumiao/models"
)

//获取全部书籍
func GetBooks(c *gin.Context) {
	var data interface {}
	data=models.GetAllBook()
	c.JSON(200, gin.H{
		"message": "success",
		"data":data,
	})
}

//获取指定书籍
func GetBook(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))


	data:=models.GetBook(id)


	c.JSON(200, gin.H{
		"message":"success",
		"data":data,
	})

}

//添加书籍
func AddBook(c *gin.Context) {
	var form models.Book
	if c.Bind(&form)==nil{
		if models.AddBook(form){
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

//删除书籍
func DeleteBook(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	models.DeleteBook(id)

	c.JSON(200, gin.H{
		"message": "success",})
}
//修改书籍信息
func EditBook(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	var form models.Book
	if c.Bind(&form)==nil{
		if models.EditBook(id,form){
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
