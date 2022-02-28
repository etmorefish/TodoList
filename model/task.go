package model

import "github.com/jinzhu/gorm"

type Task struct {
	gorm.Model
	User User `gorm:"ForeignKey:Uid"`
	Uid uint `gorm:"not null"`
	Title string `gorm:"index;not null"`
	Status string `gorm:"default:'0'"`  // default: 0 未完成 1 完成
	Content string `gorm:"type:longtext"`
	StartTime int64  
	EndTime int64
}
