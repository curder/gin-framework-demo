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
		var (
			query string
			ok    bool
		)

		query = c.Query("search")                  // 获取地址栏传递的search这个key的值，不传递则为空

		query = c.DefaultQuery("search", "curder") // 如果没有search这个key，则值curder值

		if query, ok = c.GetQuery("search"); !ok { // 获取地址栏传递的search这个key的值，不传递则设置一个默认值，跟 DefaultQuery 类似功能
			query = "curder"
		}

		c.JSON(http.StatusOK, gin.H{
			"query": query,
		})
	})

	if err = r.Run(":8080"); err != nil {
		fmt.Printf("Failed start server, error: %s\n", err.Error())
		return
	}
}
