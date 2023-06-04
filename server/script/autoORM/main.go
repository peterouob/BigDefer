package main

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name     string `json:"name"`
	UserName string `json:"user"`
	Password string `json:"password"`
	Token    string `json:"token"`
}

func (*User) TableName() string {
	return "user"
}

func main() {
	var db *gorm.DB
	var err error
	db, err = gorm.Open(mysql.Open("newuser:password@tcp(localhost:3306)/bigdefer?charset=utf8mb4&parseTime=True&loc=Local"))
	if err != nil {
		fmt.Println("connection error: ", err)
		return
	}
	db.AutoMigrate(User{})
}
