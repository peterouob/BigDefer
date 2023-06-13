package router

import (
	"bigdefer/middleware"
	"bigdefer/service/user"
	sudoService "bigdefer/service/user/sudo"
	vip2 "bigdefer/service/vip"
	"github.com/gin-gonic/gin"
)

func Router(g *gin.Engine) {
	sudo := g.Group("/sudo")
	{
		sudo.POST("/delete", sudoService.DeleteUser)
		sudo.POST("/update", sudoService.UpdateUser)
		sudo.GET("/getAllUsers", sudoService.FindAllUser)
	}
	user := g.Group("/user")
	{
		user.POST("/create", userService.CreateUser)
		user.POST("/login", userService.LoginUser)
		user.POST("/refershToken", userService.RefershToken)
	}
	user.Use(middleware.Auth())
	{
		//loginSuccess := user.Group("/s")
		//{
		//	loginSuccess.GET("/report", blackuser.Report)
		//}
		vip := user.Group("/vip")
		{
			vip.GET("/vip", vip2.IamVIP)
		}
	}
}
