package user

import "github.com/runningape/goblog/pkg/model"

func (user *User) Create() (err error) {
	if err := model.DB.Create(&user).Error; err != nil {
		return err
	}
	return nil
}
