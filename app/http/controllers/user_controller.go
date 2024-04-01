package controllers

import (
	"fmt"
	"net/http"

	"github.com/runningape/goblog/app/models/article"
	"github.com/runningape/goblog/app/models/user"
	"github.com/runningape/goblog/logger"
	"github.com/runningape/goblog/pkg/route"
	"github.com/runningape/goblog/pkg/view"
	"gorm.io/gorm"
)

type UserController struct{}

func (*UserController) Show(w http.ResponseWriter, r *http.Request) {

	id := route.GetRouteVariable("id", r)

	_user, err := user.Get(id)

	if err != nil {
		if err == gorm.ErrRecordNotFound {
			w.WriteHeader(http.StatusNotFound)
			fmt.Fprint(w, "404 用户未找到。")
		} else {
			logger.LogError(err)
			w.WriteHeader(http.StatusInternalServerError)
			fmt.Fprint(w, "500 服务器内部错误。")
		}
	} else {
		articles, err := article.GetByUserID(_user.GetStringID())

		if err != nil {
			logger.LogError(err)
			w.WriteHeader(http.StatusInternalServerError)
			fmt.Fprint(w, "500 服务器内部错误。")
		} else {
			view.Render(w, view.D{
				"Articles": articles,
			}, "articles.index", "articles._article_meta")
		}
	}
}
