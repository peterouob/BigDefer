package main

import (
	"bigdefer/config"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"status": "ok",
		})
	})
	r.Run(":80")
	fmt.Println(config.Config.GetString("mysql.dsn")) //測試config
}
