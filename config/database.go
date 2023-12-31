package config

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDb() {
	dsn := "dev:dev@tcp(127.0.0.1:3306)/pos_alta?charset=utf8mb4&parseTime=True&loc=Local"

	var err error
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	
	if err != nil {
		panic(err)
	}
}