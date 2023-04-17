package main

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gen"
	"gorm.io/gorm"
)

func main() {
	var db *gorm.DB
	var err error
	db, err = gorm.Open(mysql.Open("newuser:password@tcp(localhost:3306)/bigdefer?charset=utf8mb4&parseTime=True&loc=Local"))
	if err != nil {
		panic(fmt.Errorf("%s", "connect to mysql error"))
	}
	g := gen.NewGenerator(gen.Config{
		OutPath:      "/Users/peter/Desktop/BigDefer/orm/dal",
		ModelPkgPath: "/Users/peter/Desktop/BigDefer/orm/model",
		Mode:         gen.WithDefaultQuery | gen.WithoutContext,
	})
	g.UseDB(db)
	g.ApplyBasic(g.GenerateAllTable()...)
	g.Execute()
}
