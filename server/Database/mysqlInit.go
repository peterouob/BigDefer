package Database

import (
	"bigdefer/config"
	"bigdefer/orm/dal"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func MysqlInit() {
	db, err := gorm.Open(mysql.Open(config.Config.GetString("mysql.dsn")), &gorm.Config{})
	if err != nil {
		fmt.Println("Connect Error: ", err)
		return
	}
	dal.SetDefault(db)
	fmt.Println("Connect succeeded ... ...")
}
