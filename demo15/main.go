package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	var (
		r   *gin.Engine
		err error
	)
	r = gin.Default()

	r.LoadHTMLFiles("./index.tpl")

	r.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.tpl", gin.H{})
	})

	r.POST("/", func(c *gin.Context) {
		var (
			username string
			password string
			ok       bool
		)
		// 获取form表单通过post请求的请求主体

		// 方式一：PostForm
		username = c.PostForm("username")
		password = c.PostForm("password")

		// 方式二：DefaultPostForm
		username = c.DefaultPostForm("username", "curder") // 获取username参数的值，如果不存在username则赋予默认值
		password = c.DefaultPostForm("password", "******") // 获取password参数的值，如果不存在password则赋予默认值

		// 方式三：GetPostForm
		if username, ok = c.GetPostForm("username"); !ok {
			username = "curder"
		}
		if password, ok = c.GetPostForm("password"); !ok {
			password = "******"
		}

		c.JSON(http.StatusOK, gin.H{
			"username": username,
			"password": password,
		})

	})

	if err = r.Run(":8080"); err != nil {
		fmt.Printf("Server start error: %s\n", err.Error())
		return
	}
}
