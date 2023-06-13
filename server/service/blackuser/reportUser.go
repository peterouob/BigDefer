package blackuser

import (
	"bigdefer/utils"
	"github.com/gin-gonic/gin"
)

var BlackQueue chan string

type BlackReportRequest struct {
	FromUser string `json:"from_user"`
	Message  string `json:"message"`
	ToUser   string `json:"to_user"`
}

func ChekBlackQuestion(c *gin.Context) {
	var req BlackReportRequest
	c.ShouldBind(&req)
	select {
	case BlackQueue <- req.FromUser:
		//TODO:查詢使用者同時間檢舉別人數量
	case BlackQueue <- req.Message:
		utils.BlackReport(<-BlackQueue)
	case BlackQueue <- req.ToUser:
		//TODO:查詢使用者被同時檢舉數量
	}
}
