package main

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/gorilla/mux"
)

var router = mux.NewRouter()

func homeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<h1>Hello,这里是goblog</h1>")
}

func aboutHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "请联系："+"<a href=\"mailto:test@test.com\">runningape</a>")
}

func notFoundHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotFound)
	fmt.Fprint(w, "<h1>请求页面未找到:(</h1><p>Please Contact us</p>")
}

func articlesShowHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	fmt.Fprint(w, "文章 ID:"+id)
}

func articlesIndexHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "访问文章列表")
}

func articlesStoreHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "r.Form中title的值为：%v <br>", r.FormValue("title"))
	fmt.Fprintf(w, "r.PostForm中title的值为: %v <br>", r.PostFormValue("title"))
	fmt.Fprintf(w, "r.Form中 test 的值为：%v <br>", r.FormValue("test"))
	fmt.Fprintf(w, "r.PostForm 中 test 的值为：%v <br>", r.PostFormValue("test"))
}

func articlesCreateHandler(w http.ResponseWriter, r *http.Request) {
	html := `
		<!DOCTYPE html>
			<head>
				<title>创建文章 - 我的技术博客</title>
			</head>
			<body>
				<form action="%s?test=data" method="post">
					<p><input type="text" name="title"></p>
					<p><textarea name="body" cols="30" rows="10"></textarea>
					<p><button type="submit">提交</button></p>
				</form>
			</body>
		</html>
	`
	storeURL, _ := router.Get("articles.store").URL()
	fmt.Fprintf(w, html, storeURL)
}

func forceHTMLMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/html; charset=utf-8")
		next.ServeHTTP(w, r)
	})
}

func removeTrailingSlash(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/" {
			r.URL.Path = strings.TrimSuffix(r.URL.Path, "/")
		}
		next.ServeHTTP(w, r)
	})
}

func main() {
	router.HandleFunc("/", homeHandler).Methods("GET").Name("home")
	router.HandleFunc("/about", aboutHandler).Methods("GET").Name("about")

	router.HandleFunc("/articles/{id:[0-9]+}",
		articlesShowHandler).Methods("GET").Name("articles.show")
	router.HandleFunc("/articles",
		articlesIndexHandler).Methods("GET").Name("articles.index")
	router.HandleFunc("/articles",
		articlesStoreHandler).Methods("POST").Name("articles.store")
	router.HandleFunc("/articles/create",
		articlesCreateHandler).Methods("GET").Name("articles.create")

	router.NotFoundHandler = http.HandlerFunc(notFoundHandler)

	router.Use(forceHTMLMiddleware)

	http.ListenAndServe(":3000", removeTrailingSlash(router))
}
