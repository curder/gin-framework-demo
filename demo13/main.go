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

	r.GET("/", jsonHandle)
	r.GET("/s", structHandle)

	if err = r.Run(":8080"); err != nil {
		fmt.Printf("start server error: %s\n", err.Error())
		return
	}
}

// jsonHandle
func jsonHandle(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"code":    http.StatusOK,
		"message": "OK",
	})
}

// structHandle
func structHandle(c *gin.Context) {
	type Result struct {
		Code    int         `json:"code"`
		Message string      `json:"message"`
		Data    interface{} `json:"data"`
	}

	c.JSON(http.StatusOK, Result{
		Data: map[string]interface{}{
			"page":   1,
			"result": true,
		},
		Message: "OK!",
		Code:    http.StatusOK,
	})
}
