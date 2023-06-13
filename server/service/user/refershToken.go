package userService

import (
	H "bigdefer/service/http"
	"bigdefer/utils"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"time"
)

func RefershToken(c *gin.Context) {
	token := utils.Token{
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(7 * 24 * time.Hour).UnixMilli(),
		},
	}
	tk, err := utils.SignToken(&token)
	if err != nil {
		H.Forbidden(c, "continue token error")
	} else {
		H.SetCookieForToken(c, "new-token", tk)
		H.OK(c, "token updated successfully")
	}
}
