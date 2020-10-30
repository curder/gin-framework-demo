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
		fmt.Printf("Failed to start server, error: %s", err.Error())
	}
}

// defaultHandle
func defaultHandle(w http.ResponseWriter, r *http.Request) {
	var (
		tpl *template.Template
		err error
	)

	// 解析模版
	if tpl, err = template.ParseFiles("./index.tpl"); err != nil {
		fmt.Printf("Parse file error: %s", err.Error())
		return
	}

	var data = "Curder"

	// 渲染模版
	if err = tpl.Execute(w, data); err != nil {
		fmt.Printf("Failed to render template, error: %s", err.Error())
		return
	}
}
