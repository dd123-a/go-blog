package Controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"twqblog/Model"
)

func IndexHandler(c *gin.Context) {
	//获取get参数
	sort :=c.DefaultQuery("sort","1")//1,最新，2，最热
	cate :=c.DefaultQuery("cate","")//分类
	tag :=c.DefaultQuery("tag","")//标签
	titleMap :=make(map[string]string)

	var isHot bool

	if sort == "1"{
		isHot=false
		titleMap["firstName"]="最新发布"
		titleMap["secondName"]="热门文章"
		titleMap["secUrl"]="?sort=2"
	}else {
		isHot=true
		titleMap["firstName"]="热门文章"
		titleMap["secondName"]="热门文章"
		titleMap["secUrl"]="?sort=2"
	}

	blogModel :=Model.Blog{}
	blogList :=blogModel.GetBlogList(isHot,cate,tag)

	//右侧栏列表
	secList :=blogModel.GetSecBlogList(!isHot)
	fmt.Println(secList)

	//标签
	tagModel :=Model.Tag{}
	tagList :=tagModel.GetTagList(tag)

	//分类
	cateModel :=Model.Cate{}
	cateList :=cateModel.GetCateList(cate)

	c.HTML(http.StatusOK,"index/index.html",gin.H{
		"titleMap":titleMap,
		"blogCount":len(blogList),
		"blogList": blogList,
		"secList":secList,
		"tagList":tagList,
		"cateList":cateList,
	})
}




































