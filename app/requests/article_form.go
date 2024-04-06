package requests

import (
	"github.com/runningape/goblog/app/models/article"
	"github.com/thedevsaddam/govalidator"
)

func ValidateArticleForm(data article.Article) map[string][]string {
	rules := govalidator.MapData{
		"title": []string{"required", "min:3", "max:40"},
		"body":  []string{"required", "min:10"},
	}

	messages := govalidator.MapData{
		"title": []string{
			"required:标题为必填项。",
			"min_cn:标题长度需大于 3.",
			"max_cn:标题长度需小于 40.",
		},
		"body": []string{
			"required:文章内容为必填项。",
			"min_cn:文件长度需大于 10.",
		},
	}

	opts := govalidator.Options{
		Data:          &data,
		Rules:         rules,
		TagIdentifier: "valid",
		Messages:      messages,
	}

	return govalidator.New(opts).ValidateStruct()
}
