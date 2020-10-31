package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

type User struct {
	Username string `form:"username"`
	Password string `form:"password"`
}

func main() {
	var (
		c   *gin.Engine
		err error
	)

	c = gin.Default()

	c.GET("/", getHandle)
	c.POST("/", postHandle)

	if err = c.Run(":8080"); err != nil {
		fmt.Printf("Server Start error: %s\n", err.Error())
		return
	}
}

// getHandle
func getHandle(c *gin.Context) {
	var (
		user User
		err  error
	)
	if err = c.ShouldBind(&user); err != nil { // 绑定结构体需要传递结构体对象的指针，否则将无法修改原结构体
		fmt.Printf("bind params error: %s\n", err.Error())
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status": true,
		"data": map[string]string{
			"username": user.Username,
			"password": user.Password,
		},
	})
}

// postHandel
func postHandle(c *gin.Context) {
	var (
		user User
		err  error
	)
	if err = c.ShouldBind(&user); err != nil {
		fmt.Printf("Bind error: %s\n", err.Error())
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status": true,
		"data": map[string]string{
			"username": user.Username,
			"password": user.Password,
		},
	})
}
