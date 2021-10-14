package utils

import (
	constantname "bloggo/constantName"
	"bloggo/models"
	"errors"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func ExtractDB(c *gin.Context) (*gorm.DB, error) {
	db, exist := c.Get(constantname.DB)
	if !exist {
		return nil, errors.New("db doesn't exist")
	}
	return db.(*gorm.DB), nil
}

func ExtractUserContext(c *gin.Context) (models.User, error) {
	usr, exist := c.Get(constantname.USER_CTX)
	if !exist {
		return models.User{}, errors.New("user context doesn't exist")
	}
	return usr.(models.User), nil
}
