package main

import (
	"bigdefer/orm/dal"
	"bigdefer/orm/model"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	gorm.Open(mysql.Open("root:peter63674782@tcp(localhost:3306)/bigdefer?charset=utf8mb4&parseTime=True&loc=Local"))
	user := model.User{
		Name:     "peter",
		UserName: "defer",
		Password: "12345",
	}
	dal.User.Create(&user)
}
