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

	c.GET("/user/:name/:age", func(c *gin.Context) {
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

	if err = c.Run(":8080"); err != nil {
		fmt.Printf("Server start failed, error: %s", err.Error())
		return
	}
}
