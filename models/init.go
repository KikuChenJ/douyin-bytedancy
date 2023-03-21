package models

import (
	"github.com/hjk-cloud/douyin/define"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
)

var db = Init()

func Init() *gorm.DB {
	dsn := define.DBUserName + ":" + define.DBPassWord + "@tcp(127.0.0.1:3306)/" + define.DBName + "?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Print("gorm Init Error: ", err)
	}
	return db
}
