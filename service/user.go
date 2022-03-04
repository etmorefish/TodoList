package service

import (
	"todo-list/model"
	"todo-list/serializer"
)

type UserService struct {
	UserName string `form:"user_name" json:"user_name" binding:"required,min=3,max=15"`
	Password string `form:"password" json:"password" binding:"required,min=5,max=16"`
}

func (service *UserService) Register() serializer.Response {
	var count int
	var user model.User
	model.DB.Model(&model.User{}).Where("user_name=?", service.UserName).First(&user).Count(&count)

	if count == 1 {
		return serializer.Response{
			Status: 400,
			Msg:    "User already registered",
		}
	}
	user.UserName = service.UserName

	// TODO: 加密
	if err := user.SetPassword(service.Password); err != nil {
		return serializer.Response{
			Status: 400,
			Msg:    err.Error(),
		}
	}

	// 创建用户
	if err := model.DB.Create(&user).Error; err != nil {
		return serializer.Response{
			Status: 500,
			Msg:    "数据库操作错误",
		}
	}

	return serializer.Response{
		Status: 200,
		Msg:    "用户注册成功",
	}
}
