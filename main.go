package main

import (
	"bigdefer/config"
	"bigdefer/router"
	"fmt"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	router.Router(r)
	r.Run(":80")
	fmt.Println(config.Config.GetString("mysql.dsn")) //測試config
}
