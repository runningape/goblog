package main

import (
	"fmt"
	"net/http"
)

func handlerFunc(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	if r.URL.Path == "/" {
		fmt.Fprintf(w, "<h1>Hello,这里是goblog</h1>")
	} else if r.URL.Path == "/about" {
		fmt.Fprintf(w, "请联系："+"<a href=\"mailto:test@test.com\">runningape</a>")
	} else {
		fmt.Fprintf(w, "<h1>请求页面未找到 :(</h1>"+
			"<p>请联系我们</p>")
	}
}

func main() {
	http.HandleFunc("/", handlerFunc)
	http.ListenAndServe(":3000", nil)
}
