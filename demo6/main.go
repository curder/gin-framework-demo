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
		fmt.Printf("failed listen and server, error: %s\n", err.Error())
		return
	}
}

// defaultHandle
func defaultHandle(w http.ResponseWriter, r *http.Request) {
	var (
		tpl *template.Template
		err error
	)
	// 定义和解析模版
	if tpl, err = template.ParseFiles("./index.tpl", "./ul.tpl"); err != nil {
		fmt.Printf("Parse file error: %s\n", err.Error())
		return
	}

	var data = "Curder"

	// 渲染模版
	if err = tpl.Execute(w, data); err != nil {
		fmt.Printf("Parse template failed, error:%s\n", err.Error())
		return
	}
}
