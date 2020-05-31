package conf

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
	"yumiao/models"
	"yumiao/routers"
)

//启动服务器
func RunServer(){
	

	router := gin.Default()
	router=routers.InitRouter(router)


	models.InitModels()


	s := &http.Server{
		Addr:           ":8080",
		Handler:        router,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	s.ListenAndServe()

}