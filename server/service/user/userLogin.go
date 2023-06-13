package userService

import (
	"bigdefer/orm/dal"
	"bigdefer/orm/model"
	H "bigdefer/service/http"
	"bigdefer/utils"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"strconv"
	"time"
)

type UserLoginRequest struct {
	UserName string `json:"username"`
	PassWord string `json:"password"`
}

type UserLoginResponse struct {
	Token string `json:"token"`
}

var Userinfo chan *model.User

func LoginUser(c *gin.Context) {
	var req UserLoginRequest
	if err := c.ShouldBind(&req); err != nil {
		fmt.Println("Bind data error:", err)
		return
	}

	if req.UserName == "main" && req.PassWord == "main" {
		str, _ := strconv.Atoi(req.UserName)
		token := utils.Token{
			Uid:      strconv.Itoa(str),
			Username: req.UserName,
			StandardClaims: jwt.StandardClaims{
				ExpiresAt: time.Now().Add(7 * 24 * time.Hour).UnixMilli(),
			},
		}
		tk, _ := utils.SignToken(&token)
		H.SetCookieForToken(c, "main", tk)
		H.OK(c, "管理員你好")
	} else {
		acc, err := dal.User.Where(dal.User.UserName.Eq(req.UserName)).First()
		if err != nil {
			H.Fail(c, "Login error")
			return
		}
		if len(req.PassWord) > 0 && req.PassWord == acc.Password {
			token := utils.Token{
				Uid:      strconv.Itoa(int(acc.ID)),
				Username: acc.UserName,
				StandardClaims: jwt.StandardClaims{
					ExpiresAt: time.Now().Add(7 * 24 * time.Hour).UnixMilli(),
				},
			}
			tk, err := utils.SignToken(&token)
			if err != nil {
				H.Fail(c, "token error")
				return
			} else {
				Userinfo <- acc
				H.SetCookieForToken(c, "token", tk)
				H.OK(c, "success login")
			}
		} else {
			H.Fail(c, "Login error")
			return
		}
	}
}
