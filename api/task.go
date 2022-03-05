package api

import (
	"todo-list/pkg/util"
	"todo-list/service"

	"github.com/gin-gonic/gin"
	logging "github.com/sirupsen/logrus"
)

func CreateTask(c *gin.Context) {
	var createTask service.CreateTaskService
	claims, _ := util.ParseToken(c.GetHeader("Authorization"))
	if err := c.ShouldBind(&createTask); err == nil {
		res := createTask.Create(claims.Id)
		c.JSON(200, res)
	} else {
		logging.Error(err)
		c.JSON(400, ErrorResponse(err))
	}
}
