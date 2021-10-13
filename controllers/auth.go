package controllers

import (
	"bloggo/models"
	"encoding/json"
	"net/http"

	"github.com/gin-gonic/gin"
)

const AUTH_NAME = "authorization"
const AUTH_MAX_AGE = 100000
const AUTH_PATH = "/"
const AUTH_DOMAIN = "localhost"
const AUTH_HTTPONLY = false
const AUTH_SECURE = false

const TOKEN_TYPE_NEW = "new"
const TOKEN_TYPE_REFRESH = "refresh"

type AuthCtl struct {
}

// AuthLogin godoc
// @Tags auth
// @Summary send username and password to get JWT Token
// @Description send username and password to get JWT Token
// @Param  user body swagmodel.LoginPayload true "create an user"
// @Accept json
// @Produce  json
// @Success 200 {object} swagmodel.LoginResponse
// @Router /login [post]
func (c *AuthCtl) Login() gin.HandlerFunc {
	return func(c *gin.Context) {
		//decode login payload (username, password)
		var usr models.User
		err := json.NewDecoder(c.Request.Body).Decode(&usr)
		if err != nil {
			c.JSON(http.StatusInternalServerError, err.Error())
			return
		}
		//generate token
		token, err := usr.GenerateToken()
		if err != nil {
			c.JSON(http.StatusInternalServerError, err.Error())
			return
		}
		//send Response
		resp := models.LoginResponse{
			TokenType: TOKEN_TYPE_NEW,
			Token:     token,
		}
		c.JSON(http.StatusOK, resp)
	}
}

func (c *AuthCtl) Register() gin.HandlerFunc {
	return func(c *gin.Context) {

	}
}
