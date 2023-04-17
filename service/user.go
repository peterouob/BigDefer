package service

import (
	H "bigdefer/service/http"

	"github.com/gin-gonic/gin"
)

func GetUser(c *gin.Context) {
	H.OK(c, "success")
}
