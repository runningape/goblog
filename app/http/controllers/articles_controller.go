package controllers

import (
	"fmt"
	"html/template"
	"net/http"
	"path/filepath"
	"strconv"
	"unicode/utf8"

	"github.com/runningape/goblog/app/models/article"
	"github.com/runningape/goblog/logger"
	"github.com/runningape/goblog/pkg/route"
	"github.com/runningape/goblog/pkg/types"
	"gorm.io/gorm"
)

type ArticlesController struct {
}

type ArticlesFormData struct {
	Title, Body string
	URL         string
	Errors      map[string]string
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

func (*ArticlesController) Index(w http.ResponseWriter, r *http.Request) {
	article, err := article.GetAll()

	if err != nil {
		logger.LogError(err)
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprint(w, "500 Internal Server Error")
	} else {
		viewDir := "resources/views"
		files, err := filepath.Glob(viewDir + "/layouts/*.html")
		logger.LogError(err)
		newfiles := append(files, viewDir+"/articles/index.html")
		tmpl, err := template.ParseFiles(newfiles...)
		logger.LogError(err)
		err = tmpl.ExecuteTemplate(w, "app", article)
		logger.LogError(err)
	}
}

func (*ArticlesController) Create(w http.ResponseWriter, r *http.Request) {
	storeURL := route.Name2URL("articles.store")

	data := ArticlesFormData{
		Title:  "",
		Body:   "",
		URL:    storeURL,
		Errors: nil,
	}

	tmpl, err := template.ParseFiles("resources/views/articles/create.html")
	if err != nil {
		panic(err)
	}
	err = tmpl.Execute(w, data)
	if err != nil {
		panic(err)
	}
}

func validateArticleFormData(title, body string) map[string]string {
	errors := make(map[string]string)

	if title == "" {
		errors["title"] = "The title length cannot be empty."
	} else if utf8.RuneCountInString(title) < 3 || utf8.RuneCountInString(title) > 40 {
		errors["title"] = "The title length must be between 3 and 40 characters."
	}

	if body == "" {
		errors["body"] = "The content length cannot be empty."
	} else if utf8.RuneCountInString(body) < 10 {
		errors["body"] = "The content length must be less than 10 characters."
	}

	return errors

}

func (*ArticlesController) Store(w http.ResponseWriter, r *http.Request) {

	title := r.PostFormValue("title")
	body := r.PostFormValue("body")

	errors := validateArticleFormData(title, body)

	if len(errors) == 0 {
		_article := article.Article{
			Title: title,
			Body:  body,
		}
		_article.Create()
		if _article.ID > 0 {
			fmt.Fprint(w, "Insert successful. ID:"+strconv.FormatUint(_article.ID, 10))
		} else {
			w.WriteHeader(http.StatusInternalServerError)
			fmt.Fprint(w, "Failed to create the article, please contact the administrator!")
		}

	} else {
		storeURL := route.Name2URL("articles.store")
		data := ArticlesFormData{
			Title:  title,
			Body:   body,
			URL:    storeURL,
			Errors: errors,
		}

		tmpl, err := template.ParseFiles("resources/views/articles/create.html")
		logger.LogError(err)

		err = tmpl.Execute(w, data)
		logger.LogError(err)
	}
}

func (*ArticlesController) Edit(w http.ResponseWriter, r *http.Request) {

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
		updateURL := route.Name2URL("articles.update", "id", id)
		data := ArticlesFormData{
			Title:  article.Title,
			Body:   article.Body,
			URL:    updateURL,
			Errors: nil,
		}

		tmpl, err := template.ParseFiles("resources/views/articles/edit.html")
		logger.LogError(err)
		err = tmpl.Execute(w, data)
		logger.LogError(err)
	}
}

func (*ArticlesController) Update(w http.ResponseWriter, r *http.Request) {

	id := route.GetRouteVariable("id", r)
	_article, err := article.Get(id)

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
		title := r.PostFormValue("title")
		body := r.PostFormValue("body")

		errors := validateArticleFormData(title, body)

		if len(errors) == 0 {
			_article.Title = title
			_article.Body = body
			rowsAffected, err := _article.Update()

			if err != nil {
				logger.LogError(err)
				w.WriteHeader(http.StatusInternalServerError)
				fmt.Fprint(w, "500 Internal Server Error")
			}

			if rowsAffected > 0 {
				showURL := route.Name2URL("articles.index")
				http.Redirect(w, r, showURL, http.StatusFound)
			} else {
				fmt.Fprint(w, "You haven't made any changes.")
			}

		} else {
			updateURL := route.Name2URL("articles.update", "id", id)
			data := ArticlesFormData{
				Title:  title,
				Body:   body,
				URL:    updateURL,
				Errors: errors,
			}

			tmpl, err := template.ParseFiles("resources/views/articles/edit.html")
			logger.LogError(err)
			err = tmpl.Execute(w, data)
			logger.LogError(err)
		}
	}
}

func (*ArticlesController) Delete(w http.ResponseWriter, r *http.Request) {

	id := route.GetRouteVariable("id", r)
	_article, err := article.Get(id)

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
		rowsAffected, err := _article.Delete()

		if err != nil {
			logger.LogError(err)
			w.WriteHeader(http.StatusInternalServerError)
			fmt.Fprint(w, "500 Internal Server Error")
		} else {
			if rowsAffected > 0 {
				indexURL := route.Name2URL("articles.index")
				http.Redirect(w, r, indexURL, http.StatusFound)
			} else {
				w.WriteHeader(http.StatusNotFound)
				fmt.Fprint(w, "404 article not found")
			}
		}
	}
}
