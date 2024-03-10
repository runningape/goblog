package controllers

import (
	"fmt"
	"net/http"
	"strconv"
	"unicode/utf8"

	"github.com/runningape/goblog/app/models/article"
	"github.com/runningape/goblog/pkg/logger"
	"github.com/runningape/goblog/pkg/route"
	"github.com/runningape/goblog/pkg/view"
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
			fmt.Fprint(w, "404 Article Not Found")
		} else {
			logger.LogError(err)
			w.WriteHeader(http.StatusInternalServerError)
			fmt.Fprint(w, "500 Internal Server Error")
		}
	} else {
		view.Render(w, article, "articles.show")
	}
}

func (*ArticlesController) Index(w http.ResponseWriter, r *http.Request) {
	articles, err := article.GetAll()
	if err != nil {
		logger.LogError(err)
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprint(w, "500 Internal Server Error")
	} else {
		view.Render(w, articles, "articles.index")
	}
}

func validateArticleFormData(title, body string) map[string]string {
	errors := make(map[string]string)

	if title == "" {
		errors["title"] = "The title cannot be empty."
	} else if utf8.RuneCountInString(title) < 3 ||
		utf8.RuneCountInString(title) > 40 {
		errors["title"] = "The title length must be between 3 and 40 characters."
	}

	if body == "" {
		errors["body"] = "The content cannot be empty."
	} else if utf8.RuneCountInString(body) < 10 {
		errors["body"] = "The content length cannot be less than 10 characters"
	}

	return errors
}

func (*ArticlesController) Create(w http.ResponseWriter, r *http.Request) {
	view.Render(w, view.D{}, "articles.create", "articles._form_field")
}

func (*ArticlesController) Store(w http.ResponseWriter, r *http.Request) {

	title := r.PostFormValue("title")
	body := r.PostFormValue("body")

	errors := validateArticleFormData(title, body)

	if len(errors) == 0 {
		_article := article.Article{
			Body:  body,
			Title: title,
		}

		_article.Create()

		if _article.ID > 0 {
			fmt.Fprint(w, "Insert successful! ID:"+strconv.FormatUint(_article.ID, 10))
		} else {
			w.WriteHeader(http.StatusInternalServerError)
			fmt.Fprint(w, "Failed to create article, please contact the administrator!")
		}

	} else {
		view.Render(w, view.D{
			"Title":  title,
			"Body":   body,
			"Errors": errors,
		}, "articles.create", "articles._form_field")
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
		view.Render(w, view.D{
			"Title":   article.Title,
			"Body":    article.Body,
			"Article": article,
			"Errors":  nil,
		}, "articles.edit", "articles._form_field")
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
			fmt.Fprint(w, "500 Internal Sever Error")
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
				fmt.Fprint(w, "You didn't make and changes")
			}
		} else {
			view.Render(w, view.D{
				"Title":   title,
				"Body":    body,
				"Article": _article,
				"Errors":  errors,
			}, "articles.edit", "articles._form_field")
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
