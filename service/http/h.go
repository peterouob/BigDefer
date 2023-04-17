package H

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func OK(c *gin.Context, message string) {
	c.JSON(http.StatusOK, gin.H{
		"code":    0,
		"message": message,
	})
}

func Fail(c *gin.Context, message string) {
	c.JSON(http.StatusBadRequest, gin.H{
		"code":    -1,
		"message": message,
	})
}
