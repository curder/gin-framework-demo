package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	var (
		err error
	)

	http.HandleFunc("/", sayHello)

	if err = http.ListenAndServe(":8080", nil); err != nil {
		fmt.Printf("start http server error: %s", err.Error())
	}
}

func sayHello(w http.ResponseWriter, r *http.Request) {
	var (
		content []byte
		err     error
	)
	if content, err = ioutil.ReadFile("./main.tpl"); err != nil {
		fmt.Printf("read file failed, error: %s", err.Error())
		return
	}

	if _, err = w.Write(content); err != nil {
		fmt.Printf("server error")
		return
	}
}
