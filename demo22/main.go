package main

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

func main() {
	var (
		db  *gorm.DB
		u User
		err error
	)

	// 初始化db
	if db, err = gorm.Open("mysql", "root:root@tcp(127.0.0.1:33060)/gin_framework_demo?charset=utf8&parseTime=True&loc=Local"); err != nil {
		fmt.Printf("Open MySQL connect failed, error: %s\n", err.Error())
		return
	}

	defer db.Close() // 延迟关闭

	db.AutoMigrate(&User{}) // 自动迁移


	// 创建数据库表
	//u1 := User{
	//	Name:   "Curder",
	//	Gender: "male",
	//	Hobby:  "study",
	//}
	//db.Create(&u1)

	// 查询数据行
	db.First(&u)
	fmt.Printf("u:%#v\n", u)

	// 更新数据行
	db.Model(&u).Update("hobby", "music")

	// 删除
	db.Delete(&u)
}

type User struct {
	ID     int
	Name   string
	Gender string
	Hobby  string
}
