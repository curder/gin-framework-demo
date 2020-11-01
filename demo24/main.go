package main

import (
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"net/http"
)

var DB *gorm.DB

type Todo struct {
	ID     int    `json:"id"`
	Title  string `json:"title"`
	Status bool   `json:"status"`
}

func main() {
	var (
		r       *gin.Engine
		v1Group *gin.RouterGroup
		err     error
	)
	// 创建数据库
	// CREATE DATABASE gin_framework_demo character set utf8mb4 collate utf8mb4_general_ci;

	// 连接数据库
	initDB()
	DB.AutoMigrate(&Todo{}) // 数据库迁移

	// 初始化gin框架
	r = gin.Default()

	r.Static("/static", "./static/dist") // 静态资源
	r.LoadHTMLFiles("static/index.html") // 加载模版

	r.GET("/", index)

	v1Group = r.Group("/v1")
	{
		v1Group.GET("todo", todoList)          // Todo列表
		v1Group.POST("todo", createTodo)       // 创建todo
		v1Group.PUT("todo/:id", toggleStatus)  // 切换Todo状态
		v1Group.DELETE("todo/:id", deleteTodo) // 删除
	}

	if err = r.Run(); err != nil {
		panic(err)
	}
}

// index
func index(ctx *gin.Context) {
	ctx.HTML(http.StatusOK, "index.html", nil)
}

// todoList
func todoList(ctx *gin.Context) {
	var (
		todos []Todo
		err   error
	)
	if err = DB.Find(&todos).Error; err != nil {
		ctx.JSON(http.StatusOK, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, todos)
}

// createTodo
func createTodo(ctx *gin.Context) {
	var (
		todo Todo
		err  error
	)
	if err = ctx.BindJSON(&todo); err != nil {
		ctx.JSON(http.StatusOK, gin.H{"error": err.Error()})
		return
	}

	if err = DB.Create(&todo).Error; err != nil {
		ctx.JSON(http.StatusOK, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, todo)
}

// toggleStatus
func toggleStatus(ctx *gin.Context) {
	var (
		id   string
		todo Todo
		err  error
	)
	// 从地址栏上获取ID
	id = ctx.Params.ByName("id")

	// 查询todo是否存在
	if err = DB.Where("id = ?", id).First(&todo).Error; err != nil {
		ctx.JSON(http.StatusOK, gin.H{"error": err.Error()})
		return
	}

	// 更新todo的状态
	if err = ctx.BindJSON(&todo); err != nil {
		ctx.JSON(http.StatusOK, gin.H{"error": err.Error()})
		return
	}

	if err = DB.Save(&todo).Error; err != nil {
		ctx.JSON(http.StatusOK, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, todo)
}

// deleteTodo
func deleteTodo(ctx *gin.Context) {
	var (
		id   string
		todo Todo
		err  error
	)

	id = ctx.Params.ByName("id")

	// 查询todo是否存在
	if err = DB.Where("id = ?", id).First(&todo).Error; err != nil {
		ctx.JSON(http.StatusOK, gin.H{"error": err.Error()})
		return
	}

	if err = DB.Delete(&todo).Error; err != nil {
		ctx.JSON(http.StatusOK, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, todo)
}

func initDB() {
	var (
		err error
	)
	if DB, err = gorm.Open("mysql", "root:root@tcp(127.0.0.1:33060)/gin_framework_demo?charset=utf8mb4&parseTime=True&loc=Local"); err != nil {
		panic(err)
	}

}
