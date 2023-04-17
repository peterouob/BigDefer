package vip

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func IamVIP(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"vip": "peter",
	})
}
