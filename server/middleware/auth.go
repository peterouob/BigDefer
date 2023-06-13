package middleware

import (
	H "bigdefer/service/http"
	userService "bigdefer/service/user"
	"bigdefer/utils"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"strconv"
	"time"
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
		} else {
			u := <-userService.Userinfo
			t := time.Now().Add(7 * 24 * time.Hour).UnixMilli()
			if parse.ExpiresAt-time.Now().Unix() < t {
				token := utils.Token{
					Uid:      strconv.Itoa(int(u.ID)),
					Username: u.UserName,
					StandardClaims: jwt.StandardClaims{
						ExpiresAt: time.Now().Add(7 * 24 * time.Hour).UnixMilli(),
					},
				}
				tk, err := utils.SignToken(&token)
				if err != nil {
					H.Forbidden(c, "continue token error")
				} else {
					H.SetCookieForToken(c, "new-token", tk)
					c.Header("new-token-time", strconv.FormatInt(token.ExpiresAt, 10))
					H.OK(c, "token updated successfully")
				}
			} else {

			}
		}
		c.Set("token", parse)
		c.Next()
	}
}
