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

	r = gin.Default() // 初始化默认路由引擎

	r.GET("/", welcomeHandle) // 处理Get请求

	r.POST("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"method": "post",
		})
	})

	r.PUT("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"method": "put",
		})
	})

	r.PATCH("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"method": "patch",
		})
	})

	r.DELETE("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"method": "delete",
		})
	})

	r.OPTIONS("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"method": "options",
		})
	})


	if err = r.Run(":8080"); err != nil { // 启动Http服务
		fmt.Printf("failed to start server, error: %s", err.Error())
	}
}

// welcomeHandle 处理 /hello 请求
func welcomeHandle(c *gin.Context) {
	c.JSON(200, gin.H{
		"method": "get",
	})
}
