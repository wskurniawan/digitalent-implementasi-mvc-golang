package config

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"implementasi-mvc/app/model"
	"os"
)

func DBInit() *gorm.DB {
	db, err := gorm.Open(mysql.Open(fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local", os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_HOST"), os.Getenv("DB_PORT"), os.Getenv("DB_NAME"))), &gorm.Config{})
	if err != nil {
		panic("Failed to connect DB" + err.Error())
	}

	// automigrate here
	err = db.AutoMigrate(new(model.Account), new(model.Transaction))
	if err != nil {
		panic(err)
	}

	return db
}
