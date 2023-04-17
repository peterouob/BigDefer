package router

import (
	"bigdefer/service/user"
	"github.com/gin-gonic/gin"
)

func Router(g *gin.Engine) {
	user := g.Group("/user")
	{
		user.POST("/create", userService.CreateUser)
	}
}
