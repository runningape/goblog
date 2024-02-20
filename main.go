package main

import (
	"fmt"
	"net/http"
)

func handlerFunc(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	if r.URL.Path == "/" {
		fmt.Fprintf(w, "<h1>welcome to goblog</h1>")
	} else if r.URL.Path == "/about" {
		fmt.Fprintf(w, "about page")
	} else {
		fmt.Fprintf(w, "not found page!")
	}
}

func main() {
	http.HandleFunc("/", handlerFunc)
	http.ListenAndServe(":3000", nil)
}
