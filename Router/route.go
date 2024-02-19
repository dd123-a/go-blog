package Router

import (
	"github.com/gin-gonic/gin"
	"twqblog/Controller"
)

func SetRounter() *gin.Engine {
	r :=gin.Default()
	//告诉gin框架模板文件使用的静态文件在哪：
	r.Static("/static","static")
	//告诉gin框架在哪里找模板文件：
	r.LoadHTMLGlob("template/**/*")

	//首页
	r.GET("/",Controller.IndexHandler)

	//博客：
	//r.GET("/blog",blog.BlogTest)
	//r.GET("/blog/list",Controller.BlogList)
	r.GET("/blog/detail",Controller.BlogDetail)
	r.POST("/blog/postDiscuss",Controller.PostDiscuss)

	return r
}























