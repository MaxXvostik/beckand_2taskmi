package database

import (
	"errors"

	"exemple.com/beckand-2taskmi/model"
)

func IsExistUserBuLogin(login string) bool {
	var count = 0
	var user model.User
	GetDB().First(&user, model.User{Login: login}).Count(&count)
	return count == 1 && user.Id > 0

}

func Add(bean interface{}) error {
	if !GetDB().NewRecord(bean) {
		return errors.New("unable to greate")
	}
	return GetDB().Create(bean).Error
}
