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

	r.Static("/public", "./public") // 加载静态资源，比如css,js,images,fonts等

	r.LoadHTMLGlob("./views/*.tpl") // 加载模版

	r.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.tpl", gin.H{
			"name": "curder",
		})
	})

	if err = r.Run(":8080"); err != nil {
		fmt.Printf("Server start failed,error: %s", err.Error())
		return
	}
}
