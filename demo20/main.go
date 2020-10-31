package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	var (
		r     *gin.Engine
		apiV1 *gin.RouterGroup
		err   error
	)

	r = gin.Default()

	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"method": "GET"})
	})

	r.POST("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"method": "POST"})
	})

	r.DELETE("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"method": "DELETE"})
	})

	// 自定义404页面
	r.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusNotFound, gin.H{
			"data": "Not Found!",
		})
	})

	// 路由分组
	apiV1 = r.Group("/v1/users")
	{
		apiV1.GET("", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{"path": "/v1/users", "method": "GET"})
		})

		apiV1.POST("", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{"path": "v1/users", "method": "POST"})
		})
	}

	r.Any("/blog", func(c *gin.Context) {
		switch c.Request.Method {
		case http.MethodGet:
			c.JSON(http.StatusOK, gin.H{"status": "OK", "method": "GET", "path": "/blog"})
		case http.MethodPost:
			c.JSON(http.StatusOK, gin.H{"status": "OK", "method": "POST", "path": "/blog"})
			// ...
		default:
			c.JSON(http.StatusNotFound, gin.H{"status": "Failed", "method": c.Request.Method, "path": "/blog"})
		}
	})

	if err = r.Run(":8080"); err != nil {
		fmt.Printf("Server start error: %s", err.Error())
		return
	}
}
