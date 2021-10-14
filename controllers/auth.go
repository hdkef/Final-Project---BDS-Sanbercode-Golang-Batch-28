package controllers

import (
	constantname "bloggo/constantName"
	"bloggo/models"
	"bloggo/utils"
	"encoding/json"
	"net/http"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

type AuthCtl struct {
}

// AuthLogin godoc
// @Tags auth
// @Summary send username and password to get JWT Token
// @Description send username and password to get JWT Token via cookie
// @Param  user body swagmodel.LoginPayload true "create an user"
// @Accept json
// @Produce  json
// @Success 200 {object} swagmodel.Response
// @Router /login [post]
func (c *AuthCtl) Login() gin.HandlerFunc {
	return func(c *gin.Context) {
		//decode login payload (username, password)
		var usr models.User
		err := json.NewDecoder(c.Request.Body).Decode(&usr)
		if err != nil {
			utils.ResponseError(c, http.StatusInternalServerError, err.Error())
			return
		}
		db, err := utils.ExtractDB(c)
		if err != nil {
			utils.ResponseError(c, http.StatusInternalServerError, err.Error())
			return
		}
		//generate token
		token, err := utils.GenerateToken(db, usr)
		if err != nil {
			utils.ResponseError(c, http.StatusInternalServerError, err.Error())
			return
		}
		//send response
		c.JSON(http.StatusOK, models.LoginResponse{
			TokenType: constantname.TOKEN_TYPE_NEW,
			Token:     token,
		})
	}
}

func (c *AuthCtl) Register() gin.HandlerFunc {
	return func(c *gin.Context) {
		var usr models.User
		err := json.NewDecoder(c.Request.Body).Decode(&usr)
		if err != nil {
			utils.ResponseError(c, http.StatusInternalServerError, err.Error())
			return
		}
		db, err := utils.ExtractDB(c)
		if err != nil {
			utils.ResponseError(c, http.StatusInternalServerError, err.Error())
			return
		}
		err = usr.Post(db)
		if err != nil {
			utils.ResponseError(c, http.StatusInternalServerError, err.Error())
			return
		}
		utils.ResponseOK(c, "user created")
	}
}

// ChPwdLogin godoc
// @Tags auth
// @Summary change user password
// @Description send authorization header with password payload to change password
// @Param  user body swagmodel.ChPwd true "change password"
// @Accept json
// @Produce  json
// @Success 200 {object} swagmodel.Response
// @Router /ch-pwd [post]
func (c *AuthCtl) ChangePWD() gin.HandlerFunc {
	return func(c *gin.Context) {
		data := struct {
			Password string `json:"password"`
		}{}

		err := json.NewDecoder(c.Request.Body).Decode(&data)
		if err != nil {
			utils.ResponseError(c, http.StatusInternalServerError, err.Error())
			return
		}

		//validate passwd
		if data.Password == "" {
			utils.ResponseError(c, http.StatusInternalServerError, "no password")
			return
		}

		//get id from token
		usr, err := utils.ExtractUserContext(c)
		if err != nil {
			utils.ResponseError(c, http.StatusInternalServerError, err.Error())
			return
		}
		id := usr.ID

		db, err := utils.ExtractDB(c)
		if err != nil {
			utils.ResponseError(c, http.StatusInternalServerError, err.Error())
			return
		}

		var oldUser models.User
		err = db.Where("id = ?", id).First(&oldUser).Error
		if err != nil {
			utils.ResponseError(c, http.StatusInternalServerError, err.Error())
			return
		}

		//hash password with bcrypt
		hashedPass, err := bcrypt.GenerateFromPassword([]byte(data.Password), 5)
		if err != nil {
			utils.ResponseError(c, http.StatusInternalServerError, err.Error())
			return
		}

		res := db.Model(&oldUser).Updates(models.User{
			Password: string(hashedPass),
		})
		if res.Error != nil {
			utils.ResponseError(c, http.StatusInternalServerError, err.Error())
			return
		}
		utils.ResponseOK(c, "password changed")
	}
}
