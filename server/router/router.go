package router

import (
	"bigdefer/middleware"
	"bigdefer/service/blackuser"
	"bigdefer/service/user"
	vip2 "bigdefer/service/vip"
	"github.com/gin-gonic/gin"
)

func Router(g *gin.Engine) {
	user := g.Group("/user")
	{
		user.POST("/create", userService.CreateUser)
		user.POST("/login", userService.LoginUser)
	}
	user.Use(middleware.Auth())
	{
		loginSuccess := user.Group("/s")
		{
			loginSuccess.GET("/report", blackuser.Report)
		}
		vip := user.Group("/vip")
		{
			vip.GET("/vip", vip2.IamVIP)
		}
	}
}
