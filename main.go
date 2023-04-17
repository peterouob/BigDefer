package main

import (
	"bigdefer/Database"
	"bigdefer/router"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	go Database.MysqlInit()
	router.Router(r)
	r.Run(":80")
}
