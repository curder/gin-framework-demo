package main

import (
	"fmt"
	"html/template"
	"net/http"
)

func main() {
	var (
		err error
	)
	http.HandleFunc("/", defaultHandle)

	if err = http.ListenAndServe(":8080", nil); err != nil {
		fmt.Printf("Failed start server, error: %s", err.Error())
		return
	}
}

type User struct {
	Name   string
	Age    int
	Gender string
}

// defaultHandle
func defaultHandle(w http.ResponseWriter, r *http.Request) {
	var (
		tpl *template.Template
		err error
	)
	// 定义和解析模版
	if tpl, err = template.ParseFiles("./index.tpl"); err != nil {
		fmt.Printf("Failed parse file, error: %s", err.Error())
		return
	}

	// 渲染模版
	var u1 = User{ // 结构体
		Name:   "Curder",
		Age:    29,
		Gender: "Male",
	}
	var hobbyList = []string{
		"study",
		"music",
		"cook",
	}
	var m1 = map[string]interface{}{ // map
		"name":   "curder",
		"age":    29,
		"gender": "male",
		"hobbyList":  hobbyList,
	}
	var data = map[string]interface{}{ // 拼合结构体和map
		"u1": u1,
		"m1": m1,
	}

	if err = tpl.Execute(w, data); err != nil {
		fmt.Printf("Failed to execute, error: %s", err.Error())
		return
	}
}
