package controllers

import (
	"database/sql"
	"fmt"
	"html/template"
	"net/http"

	"github.com/runningape/goblog/logger"
	"github.com/runningape/goblog/pkg/route"
	"github.com/runningape/goblog/pkg/types"
)

type ArticlesController struct {
}

func (*ArticlesController) Show(w http.ResponseWriter, r *http.Request) {

	id := getRouteVariable("id", r)

	article, err := getArticleByID(id)

	if err != nil {
		if err == sql.ErrNoRows {
			w.WriteHeader(http.StatusNotFound)
			fmt.Fprint(w, "404 Article not found")
		} else {
			logger.LogError(err)
			w.WriteHeader(http.StatusInternalServerError)
			fmt.Fprint(w, "500 Internal Server Error")
		}
	} else {
		tmpl, err := template.New("show.html").
			Funcs(template.FuncMap{
				"RouteNameToURL": route.Name2URL,
				"Int64ToString":  types.Int64ToString,
			}).ParseFiles("resources/views/articles/show.html")

		logger.LogError(err)

		err = tmpl.Execute(w, article)
		logger.LogError(err)
	}
}
