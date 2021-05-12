package dao

import (
	"github.com/kuludi/kuludi-gin-vue/db"
	"github.com/kuludi/kuludi-gin-vue/model"
)

func Register(user *model.User) ( error) {

	if err := db.DB.Create(user).Error; err != nil {
		return  nil
	} else {
		return err
	}
}

func GetUserByPhone(phone string) (*model.User, error) {
	user := &model.User{}
	if err := db.DB.Where("phone = ?", phone).First(&user).Error; err != nil {
		return nil, err
	}
	return user, nil
}

func IsExist(phone string) bool {

	user := &model.User{}
	db.DB.Where("phone = ?", phone).First(user)
	if user.ID != 0 {
		return true
	}
	return false

}
