package sudoService

import (
	"bigdefer/orm/dal"
	"context"
	"github.com/gin-gonic/gin"
	"net/http"
)

var ctx = context.Background()

func FindAllUser(c *gin.Context) {
	u := dal.User
	user, err := u.WithContext(ctx).Find()
	if err != nil {
		panic(err)
		return
	}
	c.JSON(http.StatusOK, user)
}
