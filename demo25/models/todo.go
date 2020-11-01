package models

import "github.com/curder/gin-framework-demo/demo24/dao"

type Todo struct {
	ID     int    `json:"id"`
	Title  string `json:"title"`
	Status bool   `json:"status"`
}

// Create 创建Todo行
func Create(todo *Todo) (err error) {
	err = dao.DB.Create(todo).Error
	return
}

// Todos 查询todo列表
func Todos() (todoList []Todo, err error) {
	err = dao.DB.Find(&todoList).Error
	return
}

// Show 通过ID查询Todo
func Show(id string) (todo Todo, err error) {
	err = dao.DB.Where("id = ?", id).First(&todo).Error

	return
}

// Update 更新Todo
func Update(todo *Todo) (err error) {
	err = dao.DB.Save(&todo).Error

	return
}

// Delete 删除Todo
func Delete(id string) (err error) {
	err = dao.DB.Where("id = ?", id).Delete(Todo{}).Error

	return
}
