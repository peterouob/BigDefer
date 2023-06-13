package middleware

import (
	H "bigdefer/service/http"
	"bigdefer/utils"
	"github.com/gin-gonic/gin"
)

func Auth() gin.HandlerFunc {
	return func(c *gin.Context) {
		token, err := c.Cookie("token")
		if err != nil {
			H.Forbidden(c, "你沒有權限進入")
			return
		}
		parse, err := utils.ParseToken(token)
		if err != nil {
			H.Forbidden(c, "你沒有權限進入")
			return
		}
		c.Set("token", parse)
		c.Next()
	}
}
