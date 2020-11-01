package routers

import (
	"github.com/curder/gin-framework-demo/demo24/controller"
	"github.com/gin-gonic/gin"
)

func InitRouter() (r *gin.Engine) {
	var (
		v1Group *gin.RouterGroup
	)

	// 初始化gin框架
	r = gin.Default()

	r.Static("/static", "./static/dist")    // 静态资源
	r.LoadHTMLFiles("templates/index.html") // 加载模版

	r.GET("/", controller.IndexHandle)

	v1Group = r.Group("/v1")
	{
		v1Group.GET("todo", controller.TodoList)          // Todo列表
		v1Group.POST("todo", controller.CreateTodo)       // 创建Todo
		v1Group.PUT("todo/:id", controller.UpdateTodo)    // 更新Todo
		v1Group.DELETE("todo/:id", controller.DeleteTodo) // 删除Todo
	}

	return
}
