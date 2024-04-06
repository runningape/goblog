package article

import (
	"net/http"

	"github.com/runningape/goblog/logger"
	"github.com/runningape/goblog/pkg/model"
	"github.com/runningape/goblog/pkg/pagination"
	"github.com/runningape/goblog/pkg/route"
	"github.com/runningape/goblog/pkg/types"
)

// Get 通过 ID 获取文章
func Get(idstr string) (Article, error) {
	var article Article
	id := types.StringToUint64(idstr)
	if err := model.DB.Preload("User").First(&article, id).Error; err != nil {
		return article, err
	}

	return article, nil
}

// GetAll 获取全部文章
func GetAll(r *http.Request, perPage int) ([]Article, pagination.ViewData, error) {
	db := model.DB.Model(Article{}).Order("created_at desc")
	_pager := pagination.New(r, db, route.Name2URL("home"), perPage)
	viewData := _pager.Paging()

	var articles []Article
	_pager.Results(&articles)

	return articles, viewData, nil
}

func (article *Article) Create() (err error) {
	if err = model.DB.Create(&article).Error; err != nil {
		logger.LogError(err)
		return err
	}
	return nil
}

func (article *Article) Update() (int64, error) {
	result := model.DB.Save(&article)
	if err := result.Error; err != nil {
		logger.LogError(err)
		return 0, err
	}
	return result.RowsAffected, nil
}

func (article *Article) Delete() (int64, error) {
	result := model.DB.Delete(&article)
	if err := result.Error; err != nil {
		return 0, err
	}
	return result.RowsAffected, nil
}

func GetByUserID(uid string) ([]Article, error) {
	var articles []Article
	if err := model.DB.Where("user_id=?", uid).Preload("User").Find(&articles).Error; err != nil {
		return articles, err
	}
	return articles, nil
}
