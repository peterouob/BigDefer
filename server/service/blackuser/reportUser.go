package blackuser

import "github.com/gin-gonic/gin"

func Report(c *gin.Context) {
	name := c.Query("string")
}
