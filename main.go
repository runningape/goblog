package main

import (
	"fmt"
	"net/http"
)

func defaultHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	if r.URL.Path == "/" {
		fmt.Fprintf(w, "<h1>welcome to goblog</h1>")
	} else {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(w, "not found page!")
	}
}

func aboutHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	fmt.Fprintf(w, "about page")
}

func main() {
	router := http.NewServeMux()

	router.HandleFunc("/", defaultHandler)
	router.HandleFunc("/about", aboutHandler)

	router.HandleFunc("/articles", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case "GET":
			fmt.Fprintf(w, "访问文章列表")
		case "POST":
			fmt.Fprintf(w, "创建新的文章")
		}
	})

	http.ListenAndServe(":3000", router)
}
