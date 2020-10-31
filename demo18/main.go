package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"mime/multipart"
	"net/http"
	"path"
)

func main() {
	var (
		r   *gin.Engine
		err error
	)

	r = gin.Default()

	r.LoadHTMLFiles("./index.tpl")

	r.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.tpl", nil)
	})

	r.POST("/", func(c *gin.Context) {
		var (
			file     *multipart.FileHeader
			distName string
			err      error
		)
		// 从请求中读取文件内容
		if file, err = c.FormFile("file"); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
		}

		// 将读取的文件保存起来
		distName = path.Join("./", file.Filename)

		if err = c.SaveUploadedFile(file, distName); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
		}

		c.JSON(http.StatusOK, gin.H{
			"status":   "ok",
			"fileName": file.Filename,
		})
	})

	if err = r.Run(":8080"); err != nil {
		fmt.Printf("Server start error: %s", err.Error())
		return
	}
}
