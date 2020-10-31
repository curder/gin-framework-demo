package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	var (
		c   *gin.Engine
		err error
	)

	c = gin.Default()

	c.Static("/public", "./public") // 加载静态资源，比如css,js,images,fonts等

	c.LoadHTMLFiles("./views/index.tpl") // 加载模版

	c.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.tpl", gin.H{
			"name": "curder",
		})
	})

	if err = c.Run(":8080"); err != nil {
		fmt.Printf("server start failed, error: %s", err.Error())
		return
	}
}
