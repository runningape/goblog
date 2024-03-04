package controllers

import (
	"fmt"
	"html/template"
	"net/http"

	"github.com/runningape/goblog/app/models/article"
	"github.com/runningape/goblog/logger"
	"github.com/runningape/goblog/pkg/route"
	"github.com/runningape/goblog/pkg/types"
	"gorm.io/gorm"
)

type ArticlesController struct {
}

func (*ArticlesController) Show(w http.ResponseWriter, r *http.Request) {

	id := route.GetRouteVariable("id", r)

	article, err := article.Get(id)

	if err != nil {
		if err == gorm.ErrRecordNotFound {
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
				"Uint64ToString": types.Uint64ToString,
			}).ParseFiles("resources/views/articles/show.html")

		logger.LogError(err)

		err = tmpl.Execute(w, article)
		logger.LogError(err)
	}
}
