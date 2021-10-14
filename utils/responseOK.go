package utils

import (
	"bloggo/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func ResponseOK(c *gin.Context, msg string) {
	c.JSON(http.StatusOK, models.Response{
		Success: true,
		Msg:     msg,
	})
}
