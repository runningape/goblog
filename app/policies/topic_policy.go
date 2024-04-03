package policies

import (
	"github.com/runningape/goblog/app/models/article"
	"github.com/runningape/goblog/pkg/auth"
)

func CanModifyArticle(_article article.Article) bool {
	return auth.User().ID == _article.UserID
}
