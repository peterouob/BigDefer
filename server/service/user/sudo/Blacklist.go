package sudoService

import (
	H "bigdefer/service/http"
	"github.com/gin-gonic/gin"
	"time"
)

type BlackListRequest struct {
	UserName string    `json:"user_name"`
	Reason   string    `json:"reason"`
	Time     time.Time `json:"time"`
}
type BlackListResponse struct {
	User map[string][]*BlackListResponse `json:"user"`
}

func Blacklist(c *gin.Context) {
	req := &BlackListRequest{}
	if err := c.ShouldBind(&req); err != nil {
		H.Fail(c, "should bind userblacklist error")
		return
	}

}
