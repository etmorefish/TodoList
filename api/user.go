package api

import (
	"fmt"
	"todo-list/service"

	"github.com/gin-gonic/gin"
)

func UserRegister(c *gin.Context) {
	var userRegister service.UserService //相当于创建了一个UserRegisterService对象，调用这个对象中的Register方法。

	if err := c.ShouldBind(&userRegister); err == nil {
		res := userRegister.Register()
		fmt.Println("---1")
		c.JSON(200, res)
	} else {
		fmt.Println("---2")
		c.JSON(400, ErrorResponse(err))
	}
}

func UserLogin(c *gin.Context) {
	var userLogin service.UserService
	if err := c.ShouldBind(&userLogin); err == nil {
		res := userLogin.Login()
		c.JSON(200, res)
	} else {
		c.JSON(400, ErrorResponse(err))
	}
}
