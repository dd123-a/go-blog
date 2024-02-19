package Model

import(
	"fmt"
	"strconv"
	"strings"
	"twqblog/Dao"
)

type Blog struct{
	ID int `json:"Id"`
	ArticleId string
	Title string
	UserId int
	CateId int
	Cate Cate `gorm:"foreignKey:CateId"`
	Pv int
	Cover string
	Sketch string
	Content string
	CreateTime string
	Author Author `gorm:"foreignKey:UserId"`
	Tag []Tag `gorm:"many2many:blog_tag_relation;"`
	Discuss []Discuss `gorm:"foreignKey:BlogId"`
}

//作者
type Author struct {
	ID int
	Name string
	Avatar string
}

//标签
type Tag struct {
	ID int
	Tag string
	Count int
	IsCurr bool
}

//博客分类：
type Cate struct {
	ID int
	Cate string
	Count int
	IsCurr bool
}

//博文内容：
type Content struct {
	ID int
	BlogId int
	MdContent string
	Content string
}

//文章标签关联表：
type BlogTag struct {
	ID int
	BlogId int
	TagId int
}


func output () map[string]interface{} {
	return map[string]interface{}{
		"code": 200,
		"msg": "ok",
		"data": "123",
	}
}

// TableName 会将 Blog 的表名重写为 `blog` , 原来为：blogs
func (Blog) TableName() string {
	return "blog"
}
func (Author) TableName() string {
	return "user"
}
func (Tag) TableName() string {
	return "blog_tag"
}
func (Cate) TableName() string {
	return "blog_cate"
}
func (Content) TableName() string {
	return "blog_detail"
}
func (BlogTag) TableName() string {
	return "blog_tag_relation"
}

//获取博客列表：
func (self *Blog)GetBlogList(isHot bool,cate string,tag string) []*Blog {
	var blogList []*Blog
	var sortSql string
	if isHot { //如果是热门文章：
		sortSql = "pv desc"
	}else{
		sortSql = "id desc"
	}
	//查询语句：
	Db := Dao.Db.Preload("Author").Preload("Tag").Preload("Cate").Order(sortSql).Where("is_deleted = 0 and status = 1")
	if(cate != ""){
		Db.Where("cate_id = ?",cate)
	}
	if(tag != ""){
		var blogTag []*BlogTag
		Dao.Db.Where("tag_id = ?",tag).Find(&blogTag)
		var blogIdList []int
		for _,value := range blogTag{
			blogIdList = append(blogIdList,value.BlogId)
		}
		Db.Where("id in ?",blogIdList)
	}

	Db.Find(&blogList)
	for key,value := range blogList {
		//只截取年月日：
		blogList[key].CreateTime = (strings.Split(value.CreateTime," "))[0]
	}

	return blogList
}


//获取第二列表数据：
func (self *Blog)GetSecBlogList(isHot bool) []*Blog {
	var blogList []*Blog
	var sortSql string
	if isHot { //如果是热门文章：
		sortSql = "pv desc"
	}else{
		sortSql = "id desc"
	}
	Dao.Db.Model(&Blog{}).Where("is_deleted = 0 and status = 1").Order(sortSql).
		Limit(5).Select("id,article_id,title,pv,create_time").Find(&blogList)

	for key,value := range blogList {
		//只截取年月日：
		blogList[key].CreateTime = (strings.Split(value.CreateTime," "))[0]
	}
	return blogList
}

//获取博客标签：
func (self *Tag)GetTagList(curr string) []*Tag {
	currTagId,err := strconv.Atoi(curr)
	if err != nil {
		currTagId = 0;
	}
	var tagList []*Tag
	Dao.Db.Model(&Tag{}).
		Raw("select id,tag,(select count(*) from blog_tag_relation b where a.id = b.tag_id) as count from blog_tag as a limit 20").
		Scan(&tagList)
	for key,value := range tagList {
		if(value.ID == currTagId){
			tagList[key].IsCurr = true
		}else{
			tagList[key].IsCurr = false
		}
	}

	for _,v := range tagList{
		fmt.Println(*v)
	}
	return tagList
}

//获取博客分类：
func (self *Cate)GetCateList(curr string) []*Cate {
	currCateId,err := strconv.Atoi(curr)
	if err != nil {
		currCateId = 0;
	}
	var cateList []*Cate
	Dao.Db.Model(&Tag{}).
		Raw("select id,cate,(select count(*) from blog b where a.id = b.cate_id) as count from blog_cate as a ").
		Scan(&cateList)
	for key,value := range cateList {
		if(value.ID == currCateId){
			cateList[key].IsCurr = true
		}else{
			cateList[key].IsCurr = false
		}
	}
	return cateList
}


//获取博客详情：
func (self *Blog)GetBlogDetail(articleId string) *Blog {
	Dao.Db.Model(&Blog{}).
		Preload("Cate").Preload("Tag").Preload("Discuss").
		Where("is_deleted = 0 and status = 1 and article_id = ?",articleId).
		Take(&self)
	for key,_ := range self.Discuss {
		self.Discuss[key].Floor = key + 1
	}

	Dao.Db.Model(&Blog{}).Where("id = ?",self.ID).Update("pv",self.Pv + 1)
	return self
}

//获取博文内容：
func (self *Content)GetBlogContent(id int) string {
	Dao.Db.Model(&Content{}).
		Where("blog_id = ? and is_deleted = 0",id).Take(&self)
	//如果sql返回值有误如何处理？

	return self.MdContent
}






























