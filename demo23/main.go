package main

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

func main() {
	var (
		db   *gorm.DB
		err  error
	)

	if db, err = gorm.Open("mysql", "root:root@tcp(127.0.0.1:33060)/gin_framework_demo?charset=utf8&parseTime=True&loc=Local"); err != nil {
		panic(err)
	}

	defer db.Close()        // 延迟关闭MySQL数据库连接
	db.AutoMigrate(&User{}) // 自动迁移数据库

	// 一般查询
	var user User
	db.Debug().First(&user) // 根据主键查询第一条记录，传递一个User类型的指针变量
	fmt.Printf("user: %#v\n", user)

	db.Debug().Last(&user) // 根据逐渐查询最后一条记录
	fmt.Printf("user: %#v\n", user)

	var users []User
	db.Debug().Find(&users) // 查询所有记录
	fmt.Printf("users: %#v\n", user)

	db.Debug().First(&user, 2) // 查询指定主键ID
	fmt.Printf("user: %#v\n", user)

}

type User struct {
	gorm.Model
	Name   string
	Gender string
	Hobby  string
}
