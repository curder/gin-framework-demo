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
	http.HandleFunc("/", index)

	if err = http.ListenAndServe(":8080", nil); err != nil {
		fmt.Printf("Failed to start server, error: %s\n", err.Error())
		return
	}
}

func index(w http.ResponseWriter, r *http.Request) {
	var (
		tpl *template.Template
		err error
	)

	// 自定义模版标识符
	tpl = template.New("index.tpl").Delims("{[", "]}")

	// 定义和解析模版
	if tpl, err = tpl.ParseFiles("./index.tpl"); err != nil {
		fmt.Printf("Parse template file failed, error: %s\n", err.Error())
		return
	}

	var data = "index"

	// 渲染模版
	if err = tpl.Execute(w, data); err != nil {
		fmt.Printf("Execute template failed, error: %s\n", err.Error())
		return
	}
}
