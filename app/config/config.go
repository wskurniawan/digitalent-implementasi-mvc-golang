package config

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"implementasi-mvc/app/model"
)

var DB *gorm.DB

func DBInit() *gorm.DB {
	db, err := gorm.Open(mysql.Open("root:@/simple_bank?charset=utf8&parseTime=True&loc=Local"), &gorm.Config{})
	if err != nil {
		panic("Failed to connect DB" + err.Error())
	}

	// automigrate here
	db.AutoMigrate(new(model.Account), new(model.Transaction))

	DB = db

	return db
}
