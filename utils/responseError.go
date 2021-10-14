package utils

import (
	"bloggo/models"

	"github.com/gin-gonic/gin"
)

func ResponseError(c *gin.Context, code int, msg string) {
	c.JSON(code, models.Response{
		Success: false,
		Msg:     msg,
	})
}
