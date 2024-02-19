package Dao

import (
	"fmt"
	"gorm.io/driver/mysql"
	_ "gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	Db *gorm.DB
)

func InitMysql() (err error) {
	//dsn := "root:tang3255332@tcp(120.79.5.254)/thinkadmin5?charset=utf8mb4"
	dsn := "root:tang3255332@tcp(127.0.0.1)/twq?charset=utf8mb4"
	Db,err = gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		return
	}
	return nil
}

func Close()  {
	sqlDb,err := Db.DB()
	if err != nil {
		fmt.Println("Db close error")
		return
	}
	_ = sqlDb.Close()
}
