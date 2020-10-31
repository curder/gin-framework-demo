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

	r.GET("/", func(c *gin.Context) {
		c.Redirect(http.StatusMovedPermanently, "https://github.com/curder")
	})

	r.GET("/a", func(c *gin.Context) {
		c.Request.URL.Path = "/b"

		r.HandleContext(c)
	})

	r.GET("b", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"path": "/b",
		})
	})

	if err = r.Run(":8080"); err != nil {
		fmt.Printf("Start server error: %s\n", err.Error())
		return
	}
}
