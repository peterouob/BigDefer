package router

import (
	"bigdefer/service"

	"github.com/gin-gonic/gin"
)

func Router(r *gin.Engine) {
	user := r.Group("/user")
	{
		user.GET("/", service.GetUser)
	}
}
