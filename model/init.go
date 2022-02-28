package model

import (
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var DB *gorm.DB

func Database(connstring string) {
	fmt.Println("connstring: ", connstring)
	db, err := gorm.Open("mysql", connstring)
	if err != nil {
		fmt.Println("Database connect error", err)
	}
	fmt.Println("Database connection established!")

	db.LogMode(true)
	if gin.Mode() == "release" {
		db.LogMode(false) // 发行版 不打印日志
	}
	db.SingularTable(true) //表名不加s
	db.DB().SetMaxIdleConns(20)
	db.DB().SetMaxOpenConns(100)
	db.DB().SetConnMaxLifetime(time.Second * 30)

	DB = db // 复制给全局db

}
