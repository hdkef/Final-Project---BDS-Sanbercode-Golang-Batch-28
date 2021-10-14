package middleware

import (
	constantname "bloggo/constantName"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func DBMiddleware(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Set(constantname.DB, db)
		c.Next()
	}
}
