package article

import (
	"strconv"

	"github.com/runningape/goblog/app/models"
	"github.com/runningape/goblog/pkg/route"
)

type Article struct {
	models.BaseModel
	Title string
	Body  string
}

func (article Article) Link() string {
	return route.Name2URL("articles.show", "id", strconv.FormatUint(article.ID, 10))
}
