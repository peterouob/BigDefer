package userService

import (
	"bigdefer/orm/dal"
	"bigdefer/orm/model"
	"github.com/gin-gonic/gin"
	"net/http"
)

type CreateUserRequest struct {
	Name     string `json:"name"`
	UserName string `json:"username"`
	PassWord string `json:"password"`
}

type CreateUserResponse struct {
	Message string `json:"message"`
	Token   string `json:"token"`
}

func CreateUser(c *gin.Context) {
	//var req CreateUserRequest
	//if err := c.ShouldBind(&req); err != nil {
	//	fmt.Println("Bind Data error", err)
	//}

	acc := model.User{
		Name:     "peter",
		UserName: "defer",
		Password: "123445",
	}
	dal.User.Create(&acc)
	c.JSON(http.StatusOK, "success")
}
