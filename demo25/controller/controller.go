package controller

import (
	"github.com/curder/gin-framework-demo/demo24/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

// IndexHandle
func IndexHandle(ctx *gin.Context) {
	ctx.HTML(http.StatusOK, "index.html", nil)
}

// createTodo
func CreateTodo(ctx *gin.Context) {
	var (
		todo models.Todo
		err  error
	)
	if err = ctx.BindJSON(&todo); err != nil {
		ctx.JSON(http.StatusOK, gin.H{"error": err.Error()})
		return
	}

	if err = models.Create(&todo); err != nil {
		ctx.JSON(http.StatusOK, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, todo)
}

// TodoList
func TodoList(ctx *gin.Context) {
	var (
		todos []models.Todo
		err   error
	)
	if todos, err = models.Todos(); err != nil {
		ctx.JSON(http.StatusOK, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, todos)
}

// UpdateTodo
func UpdateTodo(ctx *gin.Context) {
	var (
		id   string
		todo models.Todo
		err  error
	)
	// 从地址栏上获取ID
	id = ctx.Params.ByName("id")

	// 查询todo是否存在
	if todo, err = models.Show(id); err != nil {
		ctx.JSON(http.StatusOK, gin.H{"error": err.Error()})
		return
	}

	// 更新todo的状态
	if err = ctx.BindJSON(&todo); err != nil {
		ctx.JSON(http.StatusOK, gin.H{"error": err.Error()})
		return
	}

	if err = models.Update(&todo); err != nil {
		ctx.JSON(http.StatusOK, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, todo)
}

// DeleteTodo
func DeleteTodo(ctx *gin.Context) {
	var (
		id   string
		todo models.Todo
		err  error
	)

	id = ctx.Params.ByName("id")

	// 查询todo是否存在
	if todo, err = models.Show(id); err != nil {
		ctx.JSON(http.StatusOK, gin.H{"error": err.Error()})
		return
	}

	if err = models.Delete(id); err != nil {
		ctx.JSON(http.StatusOK, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, todo)
}
