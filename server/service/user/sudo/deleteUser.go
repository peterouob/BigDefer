package sudoService

import (
	"bigdefer/orm/dal"
	H "bigdefer/service/http"
	"fmt"
	"github.com/gin-gonic/gin"
	"strconv"
)

type DeleteRequest struct {
	Id string `json:"id"`
}

func DeleteUser(c *gin.Context) {
	req := &DeleteRequest{}
	if err := c.ShouldBind(req); err != nil {
		fmt.Println("bind delete request error:", err)
		return
	}
	u := dal.User
	id, _ := strconv.Atoi(req.Id)

	_, err := u.WithContext(ctx).Where(u.ID.Eq(int64(id))).Delete()
	if err != nil {
		fmt.Println("delete user error:", err)
		return
	}
	H.OK(c, "success delete")
	return
}
