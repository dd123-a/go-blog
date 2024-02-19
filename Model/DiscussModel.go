package Model

import (
	"twqblog/Dao"
)

type Discuss struct {
	ID int
	BlogId int
	Floor int `from`
	Name string
	Content string
	CreateTime string
}

func (Discuss) TableName() string {
	return "blog_discuss"
}

func (self *Discuss) PostDiscuss() bool {
	Dao.Db.Create(&self);
	return true;
}
