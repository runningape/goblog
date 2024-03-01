package controllers

import (
	"fmt"
	"net/http"
)

type PagesController struct {
}

func (*PagesController) Home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "<h1>Welcome to goblog!</h1>")
}

func (*PagesController) About(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "<h1>About Page!</h1>")
}

func (*PagesController) NotFound(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "<h1>Page not found</h1>")
}
