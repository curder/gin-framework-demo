package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"html/template"
)

func main() {
	var (
		r   *gin.Engine
		err error
	)

	r = gin.Default()

	// 在解析模版之前定义自定义的函数
	r.SetFuncMap(template.FuncMap{
		"safe": func(str string) template.HTML {
			return template.HTML(str)
		},
	})

	r.LoadHTMLGlob("./views/*.tpl") // 解析模版

	r.GET("/", func(c *gin.Context) {
		c.HTML(200, "index.tpl", gin.H{
			"name":   "curder",
			"unsafe": "<a href='https://github.com'>GitHub</a>",
		})
	})

	if err = r.Run(":8080"); err != nil {
		fmt.Printf("Failed to start server, error: %s\n", err.Error())
		return
	}

}
