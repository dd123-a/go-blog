package Controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"html"
	"html/template"
	"net/http"
	"strconv"
	"time"
	"twqblog/Model"
)

//博客列表
func BlogList(c *gin.Context) {
	blog:=Model.Blog{}
	blogList := blog.GetBlogList(false,"","")
	c.HTML(http.StatusOK,"blog/index.html",gin.H{
		"list":blogList,
		"test":"hello world",
	})
}

//博客详情
func BlogDetail(c *gin.Context)  {
	//获取get参数
	id :=c.DefaultQuery("id","0")
	blogModel :=Model.Blog{}
	blog:=blogModel.GetBlogDetail(id)

	//博客详情
	blogContentModel :=Model.Content{}
	blog.Content=blogContentModel.GetBlogContent(blog.ID)
	blog.Content=html.UnescapeString(blog.Content)

	//标签
	tagModel :=Model.Tag{}
	tagList :=tagModel.GetTagList("")
	//分类
	cateModel :=Model.Cate{}
	cateList :=cateModel.GetCateList("")

	//c.JSON(http.StatusOK,blog)
	c.HTML(http.StatusOK,"blog/detail.html",gin.H{
		"blog":blog,
		"tagList":tagList,
		"cateList":cateList,
		"content":template.HTML(blog.Content),
	})
}

//博客评论
func PostDiscuss(c *gin.Context) {
	discuss :=Model.Discuss{}
	//获取post参数
	//从 POST 请求的表单数据中获取名为 "blogId" 的参数值，并将其转换为整型。
	blogIdStr :=c.DefaultPostForm("blogId","0")
	blogId,err :=strconv.Atoi(blogIdStr)
	if err!=nil{
		fmt.Println(err)
		return
	}
	discuss.BlogId=blogId
	discuss.Name=c.DefaultPostForm("name","游客")
	discuss.Content=c.DefaultPostForm("content","")
	discuss.CreateTime=time.Unix(time.Now().Unix(),0).Format("2006-01-02 15:04:05")
	res :=discuss.PostDiscuss()

	if res {
		c.JSON(http.StatusOK,gin.H{
			"code":200,
			"msg":"评论成功",
			"data":discuss,
		})
	}else {
		c.JSON(http.StatusOK,gin.H{
			"code":500,
			"msg":"服务器繁忙",
		})
	}
}






























