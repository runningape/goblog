package requests

import (
	"github.com/runningape/goblog/app/models/user"

	"github.com/thedevsaddam/govalidator"
)

func ValidateRegistrationForm(data user.User) map[string][]string {

	rules := govalidator.MapData{
		"name":             []string{"required", "alpha_num", "between:3,20", "not_exists:users,name"},
		"email":            []string{"required", "min:4", "max:30", "email", "not_exists:users,email"},
		"password":         []string{"required", "min:6"},
		"password_confirm": []string{"required"},
	}

	messages := govalidator.MapData{
		"name": []string{
			"required:用户名为必填项。",
			"alpha_num:格式错误，只允许字母和数字。",
			"between:用户名长度在 3-20 字符之间。",
		},
		"email": []string{
			"required:Email 为必填项。",
			"min: Email 长度需大于4字符。",
			"max: Email 长度需要小于 30 字符。",
			"email:Email 格式不正确。",
		},
		"password": []string{
			"required:密码必须填写。",
			"min:密码的长度不能小于6字符",
		},
		"password_confirm": []string{
			"required:重复密码必须填写。",
		},
	}

	opts := govalidator.Options{
		Data:          &data,
		Rules:         rules,
		TagIdentifier: "valid",
		Messages:      messages,
	}

	errs := govalidator.New(opts).ValidateStruct()

	if data.Password != data.PasswordConfirm {
		errs["password_confirm"] = append(errs["password_confirm"], "两次输入的密码不匹配！")
	}

	return errs
}
