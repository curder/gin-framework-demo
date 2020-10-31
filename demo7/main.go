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
	http.HandleFunc("/about", about)

	if err = http.ListenAndServe(":8080", nil); err != nil {
		fmt.Printf("Fialed to listen and serve, error: %s", err.Error())
		return
	}
}

// index
func index(w http.ResponseWriter, r *http.Request) {
	var (
		tpl *template.Template
		err error
	)
	// 定义和渲染模版时，父模板和子模板的顺序不能乱，父在前，子在后
	if tpl, err = template.ParseFiles("./views/master.tpl", "./views/index.tpl"); err != nil {
		fmt.Printf("Parse file error: %s\n", err.Error())
		return
	}

	var data = "index"

	// 解析模版时使用 ExecuteTemplate 函数，需要制定要被渲染的模板名称
	if err = tpl.ExecuteTemplate(w, "index.tpl", data); err != nil {
		fmt.Printf("Render file error: %s\n", err.Error())
		return
	}
}

// about
func about(w http.ResponseWriter, r *http.Request) {
	var (
		tpl *template.Template
		err error
	)
	// 定义和渲染模版时，父模板和子模板的顺序不能乱，父在前，子在后
	if tpl, err = template.ParseFiles("./views/master.tpl", "./views/about.tpl"); err != nil {
		fmt.Printf("Parse file error: %s\n", err.Error())
		return
	}

	var data = "about"

	// 解析模版时使用 ExecuteTemplate 函数，需要制定要被渲染的模板名称
	if err = tpl.ExecuteTemplate(w, "about.tpl", data); err != nil {
		fmt.Printf("Render file error: %s\n", err.Error())
		return
	}
}
