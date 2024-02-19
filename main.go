package main

import (
	"fmt"
	"twqblog/Dao"
	"twqblog/Router"
)

func main() {
	//连接数据库
	err :=Dao.InitMysql()
	if err!=nil{
		fmt.Printf("init mysql err,err:#{err}\n")
		return
	}
	//即无论函数是正常结束还是发生 panic 异常结束，
	//在函数退出时都会按照 defer 语句注册的顺序逆序执行对应的延迟函数。
	defer Dao.Close()

	//2，模型绑定：传入一个结构体的内存地址，将数据库表与一个结构体绑定
	//Dao.Db.AutoMigrate(&Model.Blog{})

	//路由注册
	r :=Router.SetRounter()
	if err:=r.Run(":8010"); err!=nil{
		fmt.Printf("gin run err : #{err}\n")
	}

}