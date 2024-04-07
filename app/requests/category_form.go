package requests

import (
	"github.com/runningape/goblog/app/models/category"
	"github.com/thedevsaddam/govalidator"
)

func ValidateCategoryForm(data category.Category) map[string][]string {
	rules := govalidator.MapData{
		"name": []string{"required", "min_cn:2", "max_cn:8", "not_exists:categories,name"},
	}
	messages := govalidator.MapData{
		"name": []string{
			"required:分类名称为必填项。",
			"min_cn:分类名称长度至少需要 2 字符",
			"max_cn:分类名称长度不能超过 8 字符",
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
