package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

func main() {
	var (
		r   *gin.Engine
		err error
	)

	r = gin.Default()

	r.Use(m1, m2) // 注册全局中间件

	r.GET("/", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"message": "OK",
		})
	})

	if err = r.Run(":8080"); err != nil {
		fmt.Printf("Start server error: %s", err.Error())
		return
	}
}

// m1
func m1(ctx *gin.Context) {
	var (
		startTime time.Time
		useTime   time.Duration
	)

	fmt.Println("m1 in...")
	startTime = time.Now() // 程序开始时间

	ctx.Next() // 执行后续的handleFunc

	useTime = time.Since(startTime) // 计算前后时间差
	fmt.Printf("use time: %s\n", useTime.String())
	fmt.Println("m1 out")
}

func m2(ctx *gin.Context)  {
	fmt.Println("m2 in...")
	ctx.Next()
	fmt.Println("m2 out")
}
