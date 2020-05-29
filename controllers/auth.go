package controllers

import (
	"github.com/astaxie/beego/validation"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"strconv"
	"yumiao/app/e"
	"yumiao/app/jwt"
	"yumiao/models"
)

type auth struct {
	Username string `valid:"Required; MaxSize(50)"`
	Password string `valid:"Required; MaxSize(50)"`
}

func GetAuth(c *gin.Context) {
	username := c.Query("username")
	password := c.Query("password")

	valid := validation.Validation{}
	a := auth{Username: username, Password: password}
	ok, _ := valid.Valid(&a)

	data := make(map[string]interface{})
	code := e.INVALID_PARAMS
	if ok {
		isExist := models.CheckAuth(username, password)
		if isExist {
			token, err := jwt.GenerateToken(username, password)
			if err != nil {
				code = e.ERROR_AUTH_TOKEN
			} else {
				data["token"] = token

				code = e.SUCCESS
			}

		} else {
			code = e.ERROR_AUTH
		}
	} else {
		for _, err := range valid.Errors {
			log.Println(err.Key, err.Message)
		}
	}

	c.JSON(http.StatusOK, gin.H{
		"code" : code,
		"msg" : e.GetMsg(code),
		"data" : data,
	})
}

func Register(c *gin.Context){
	var form models.User
	if c.Bind(&form)==nil{
		if models.Register(form){
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

func Login(c *gin.Context){
	var form models.User
	if c.Bind(&form)==nil{
		if models.Login(form){

			//第一个参数为 cookie 名；
			//第二个参数为 cookie 值；
			//第三个参数为 cookie 有效时长，当 cookie 存在的时间超过设定时间时，cookie 就会失效，它就不再是我们有效的 cookie；
			//第四个参数为 cookie 所在的目录；第
			//五个为所在域，表示我们的 cookie 作用范围；
			//第六个表示是否只能通过 https 访问；
			//第七个表示 cookie 是否可以通过 js代码进行操作。
			c.SetCookie("user_cookie", strconv.Itoa(int(form.ID)), 1000, "/", "localhost", false, true)

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

func Logout(c *gin.Context){
	c.SetCookie("site_cookie", "cookievalue", -1, "/", "localhost", false, true)
	c.JSON(200, gin.H{
		"message": "success",
	})
}

