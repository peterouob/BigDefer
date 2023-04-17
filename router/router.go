package router

import (
	H "bigdefer/service/http"
	"bigdefer/service/user"
	vip2 "bigdefer/service/vip"
	"github.com/gin-gonic/gin"
)

func Router(g *gin.Engine) {
	user := g.Group("/user")
	{
		user.POST("/create", userService.CreateUser)
	}
	user.Use(H.Auth())
	vip := user.Group("/vip")
	{
		vip.GET("/vip", vip2.IamVIP)
	}
}
