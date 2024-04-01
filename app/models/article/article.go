package article

import (
	"strconv"

	"github.com/runningape/goblog/app/models/user"

	"github.com/runningape/goblog/app/models"
	"github.com/runningape/goblog/pkg/route"
)

type Article struct {
	models.BaseModel
	Title  string `gorm:"type:varchar(255);not null;" valid:"title"`
	Body   string `gorm:"type:longtext;not null;" valid:"body"`
	UserID uint64 `gorm:"not null;index"`
	User   user.User
}

func (article Article) Link() string {
	return route.Name2URL("articles.show", "id", strconv.FormatUint(article.ID, 10))
}

// CreatedAtDate 创建日期
func (article Article) CreatedAtDate() string {
	return article.CreatedAt.Format("2006-01-02")
}
