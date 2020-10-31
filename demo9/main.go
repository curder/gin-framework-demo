package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

func main() {
	var (
		r   *gin.Engine
		err error
	)

	// 初始化默认路由引擎
	r = gin.Default()

	// 加载模版目录
	r.LoadHTMLGlob("./views/*.tpl")

	r.GET("/", func(c *gin.Context) {
		// 渲染模版
		c.HTML(200, "index.tpl", gin.H{
			"name": "curder",
		})
	})

	if err = r.Run(":8080"); err != nil {
		fmt.Printf("Faield to run server, error: %s", err.Error())
	}
}
