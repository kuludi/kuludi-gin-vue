package dto

import "github.com/kuludi/kuludi-gin-vue/model"

type UserDto struct {
	Name  string `json:"name"`
	Phone string `json:"phone"`
}

func UserToDto(user *model.User) UserDto {
	return UserDto{
		Name:  user.Name,
		Phone: user.Phone,
	}
}
