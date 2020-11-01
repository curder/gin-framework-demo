package main

import (
	"github.com/curder/gin-framework-demo/demo24/dao"
	"github.com/curder/gin-framework-demo/demo24/models"
	"github.com/curder/gin-framework-demo/demo24/routers"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	var (
		r   *gin.Engine
		err error
	)

	// 创建数据库
	// CREATE DATABASE gin_framework_demo character set utf8mb4 collate utf8mb4_general_ci;

	if err = dao.InitDB(); err != nil { // 连接数据库
		panic(err)
	}

	defer dao.Close()                  // 延迟关闭数据库连接
	dao.DB.AutoMigrate(&models.Todo{}) // 数据库迁移

	r = routers.InitRouter() // 注册路由

	if err = r.Run(); err != nil { // 启动http服务
		panic(err)
	}
}
