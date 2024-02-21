package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func homeHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	fmt.Fprintf(w, "<h1>welcome to goblog</h1>")
}

func aboutHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	fmt.Fprintf(w, "about page")
}

func notFoundHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	w.WriteHeader(http.StatusNotFound)
	fmt.Fprint(w, "<h1>您访问的页面不存在。。。</h1><p>Please contact me</p>")
}

func articlesShowHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	fmt.Fprint(w, "文章 ID："+id)
}

func articlesIndexHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "文章列表页面")
}

func articlesStoreHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "创建新的文章")
}

func main() {
	router := mux.NewRouter()

	router.HandleFunc("/", homeHandler).Methods("GET").Name("home")
	router.HandleFunc("/about", aboutHandler).Methods("GET").Name("about")

	router.HandleFunc("/articles/{id:[0-9]+}",
		articlesShowHandler).Methods("GET").Name("articles.show")
	router.HandleFunc("/articles",
		articlesIndexHandler).Methods("GET").Name("articles.index")
	router.HandleFunc("/articles",
		articlesStoreHandler).Methods("POST").Name("articles.store")

	router.NotFoundHandler = http.HandlerFunc(notFoundHandler)

	homeURL, _ := router.Get("home").URL()
	fmt.Println("homeURL:", homeURL)
	articlesURL, _ := router.Get("articles.show").URL("id", "888")
	fmt.Println("articlesURL:", articlesURL)

	http.ListenAndServe(":3000", router)
}
