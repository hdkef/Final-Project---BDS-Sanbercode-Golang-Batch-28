package controllers

import (
	"bloggo/models"
	"bloggo/utils"
	"encoding/json"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	"github.com/joho/godotenv"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

const TOKEN_TYPE_NEW = "new"
const TOKEN_TYPE_REFRESH = "refresh"
const TOKEN_EXP_IN_HOUR = 12

func init() {
	godotenv.Load()
}

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
		token, err := GenerateToken(db, usr)
		if err != nil {
			utils.ResponseError(c, http.StatusInternalServerError, err.Error())
			return
		}
		//send response
		c.JSON(http.StatusOK, models.LoginResponse{
			TokenType: TOKEN_TYPE_NEW,
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

func (c *AuthCtl) ChangePWD() gin.HandlerFunc {
	return func(c *gin.Context) {

	}
}

var SECRET string = os.Getenv("SECRET")

func GenerateToken(db *gorm.DB, userFromPayload models.User) (string, error) {
	//find user detail in db
	userFromDB, err := userFromPayload.GetByUsername(db)
	if err != nil {
		return "", err
	}
	//compare db pass with payload pass
	err = bcrypt.CompareHashAndPassword([]byte(userFromDB.Password), []byte(userFromPayload.Password))
	if err != nil {
		return "", err
	}
	//generate token
	claims := jwt.MapClaims{}
	claims["id"] = userFromDB.ID
	claims["role"] = userFromDB.Role
	claims["exp"] = time.Now().Add(TOKEN_EXP_IN_HOUR * time.Hour).Unix()
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return token.SignedString([]byte(SECRET))
}
