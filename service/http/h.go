package H

import (
	"bigdefer/config"
	"bigdefer/utils"
	"github.com/gin-gonic/gin"
	"net/http"
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

func Forbidden(c *gin.Context, message string) {
	c.AbortWithStatusJSON(http.StatusForbidden, gin.H{
		"message": message,
	})
}

func Auth() gin.HandlerFunc {
	return func(c *gin.Context) {
		token, err := c.Cookie("token")
		if err != nil {
			Forbidden(c, "你沒有權限進入")
			return
		}
		parse, err := utils.ParseToken(token)
		if err != nil {
			Forbidden(c, "你沒有權限進入")
			return
		}
		c.Set("token", parse)
	}
}

func SetCookieForToken(c *gin.Context, name, value string) {
	c.SetCookie(name, value, 365*3600, "/", config.Config.GetString("server.host"), false, true)
}

func RemoveCookie(c *gin.Context, key string) {
	c.SetCookie(key, "", -1, "", config.Config.GetString("server.host"), false, true)
}
