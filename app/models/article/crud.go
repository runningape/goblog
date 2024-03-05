package article

import (
	"github.com/runningape/goblog/pkg/model"
	"github.com/runningape/goblog/pkg/types"
)

func Get(idstr string) (Article, error) {
	var article Article
	id := types.StringToUint64(idstr)
	if err := model.DB.First(&article, id).Error; err != nil {
		return article, err
	}
	return article, nil
}

func GetAll() ([]Article, error) {
	var article []Article
	if err := model.DB.Find(&article).Error; err != nil {
		return article, err
	}
	return article, nil
}
