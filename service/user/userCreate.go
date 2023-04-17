package userService

import (
	"bigdefer/orm/dal"
	"bigdefer/orm/model"
	H "bigdefer/service/http"
	"bigdefer/utils"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"strconv"
	"time"
)

type CreateUserRequest struct {
	Name       string `json:"name"`
	UserName   string `json:"username"`
	PassWord   string `json:"password"`
	RepassWord string `json:"repassword"`
	Token      string `json:"token"`
}

type CreateUserResponse struct {
	Message string `json:"message"`
	Token   string `json:"token"`
}

func CreateUser(c *gin.Context) {
	var req CreateUserRequest
	if err := c.ShouldBind(&req); err != nil {
		fmt.Println("Bind data error:", err)
		return
	}
	_, err := dal.User.Where(dal.User.UserName.Eq(req.UserName)).First()
	if err != gorm.ErrRecordNotFound {
		H.Fail(c, "This username is exist")
		return
	} else {
		fmt.Println("Hello!" + req.UserName)
	}
	if req.PassWord != req.RepassWord || req.PassWord == "" {
		H.Fail(c, "Passwords have some problem")
		return
	}
	acc := model.User{
		Name:     req.Name,
		UserName: req.UserName,
		Password: req.PassWord,
	}
	err = dal.User.Create(&acc)
	if err != nil {
		H.Fail(c, "Create failed")
		return
	}
	token := utils.Token{
		Uid:      strconv.Itoa(int(acc.ID)),
		Username: acc.UserName,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(7 * 24 * time.Hour).UnixMilli(),
		},
	}
	signToken, err := utils.SignToken(&token)
	if err != nil {
		fmt.Println("sign token error" + signToken)
		return
	} else {
		H.SetCookieForToken(c, "token", signToken)
		H.OK(c, "this is ur token"+signToken)
	}
}
