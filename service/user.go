package service

import (
	"todo-list/model"
	"todo-list/pkg/util"
	"todo-list/serializer"

	"github.com/jinzhu/gorm"
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

func (service *UserService) Login() serializer.Response {
	var user model.User
	//如果查询不到，返回相应的错误
	if err := model.DB.Where("user_name=?", service.UserName).First(&user).Error; err != nil {
		if gorm.IsRecordNotFoundError(err) {
			return serializer.Response{
				Status: 400,
				Msg:    "用户不存在，请先注册",
			}
		}
		//如果不是因为用户不存在，而是其他因素导致错误
		return serializer.Response{
			Status: 500,
			Msg:    "数据库错误",
		}
	}
	//如果查到，检查密码
	if user.CheckPassword(service.Password) == false {
		return serializer.Response{
			Status: 400,
			Msg:    "密码错误",
		}
	}

	// 创建token，为了验证用户身份
	token, err := util.GenerateToken(user.ID, service.UserName, service.Password)
	if err != nil {
		return serializer.Response{
			Status: 500,
			Msg:    "Token签发错误",
		}
	}
	return serializer.Response{
		Status: 200,
		Data: serializer.TokenData{
			User: serializer.BUildUser(user), Token: token,
		},
		Msg: "登陆成功",
	}
}
