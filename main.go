package main

import (
	"fmt"
	"html/template"
	"net/http"
)

func sayHello(w http.ResponseWriter, r *http.Request) {
	// 解析模板
	t, err := template.ParseFiles("./hello.tmpl")
	if err != nil {
		fmt.Printf("Parse template failed, err:%v", err)
		return
	}
	// 渲染模板
	err = t.Execute(w, "何大宝")
	if err != nil {
		fmt.Printf("Render template failed, err:%v", err)
		return
	}
}

func main() {
	http.HandleFunc("/", sayHello)
	err := http.ListenAndServe(":9000", nil)
	if err != nil {
		fmt.Printf("HTTP server start failed, err:%v", err)
		return
	}
}
