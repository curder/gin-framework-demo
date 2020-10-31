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
		fmt.Printf("Server start error: %s", err.Error())
		return
	}
}

// defaultHandle
func defaultHandle(w http.ResponseWriter, r *http.Request) {
	var (
		tpl        *template.Template
		customFunc func(string) (string, error)
		err        error
	)

	// 自定义函数
	customFunc = func(name string) (string, error) {
		return name + " From custom function", nil
	}

	// 定义模版
	tpl = template.New("index.tpl") // 创建自定义模版对象

	// 在解析模版之前注册自定义函数
	tpl.Funcs(template.FuncMap{
		"customFunc": customFunc,
	})

	// 解析模版
	if tpl, err = tpl.ParseFiles("./index.tpl"); err != nil {
		fmt.Printf("Parse file failed, error: %s", err.Error())
		return
	}

	// 渲染模版
	var data = "Curder"
	if err = tpl.Execute(w, data); err != nil {
		fmt.Printf("Render template error: %s", err.Error())
		return
	}

}
