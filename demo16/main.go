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

	r.GET("/user/:name/:age", func(c *gin.Context) {
		var (
			name string
			age  string
		)

		name = c.Param("name")
		age = c.Param("age")

		c.JSON(http.StatusOK, gin.H{
			"name": name,
			"age":  age,
		})
	})

	if err = r.Run(":8080"); err != nil {
		fmt.Printf("Server start failed, error: %s", err.Error())
		return
	}
}
