package sudoService

import (
	"bigdefer/orm/dal"
	"bigdefer/orm/model"
	H "bigdefer/service/http"
	"fmt"
	"github.com/gin-gonic/gin"
	"strconv"
)

type UpdateRequest struct {
	Id       string `json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
}

func UpdateUser(c *gin.Context) {
	req := &UpdateRequest{}
	err := c.ShouldBind(req)
	if err != nil {
		fmt.Println("bind user error:", err)
		return
	}
	u := dal.User
	id, _ := strconv.Atoi(req.Id)
	_, err = u.WithContext(ctx).Where(u.ID.Eq(int64(id))).Updates(model.User{
		Name:     req.Username,
		Password: req.Password,
	})
	if err != nil {
		fmt.Println("update user failed:", err)
		return
	}
	H.OK(c, "success update")
	return
}
