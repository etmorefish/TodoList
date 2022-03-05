package middleware

import (
	"time"
	"todo-list/pkg/util"

	"github.com/gin-gonic/gin"
)

func JWT() gin.HandlerFunc {
	return func(c *gin.Context) {
		code := 200
		// var data interface{}
		token := c.GetHeader("Authorization")
		if token == "" {
			code = 404
		} else {
			claims, err := util.ParseToken(token)
			if err != nil {
				code = 403 //token error
			} else if time.Now().Unix() > claims.ExpiresAt {
				code = 401 //token expireTime inviluve
			}
		}
		if code != 200 {
			c.JSON(200, gin.H{
				"status": code,
				"msg":    "token parse error",
			})
			c.Abort()
			return
		}
		c.Next()
	}
}
