package middleware

import (
	constantname "bloggo/constantName"
	"bloggo/models"
	"bloggo/utils"
	"errors"
	"fmt"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	"github.com/joho/godotenv"
)

func init() {
	godotenv.Load()
}

var SECRET string = os.Getenv("SECRET")

func JWTMiddleware(c *gin.Context) {
	token := c.GetHeader(constantname.AUTHORIZATION_HEADER)
	if token == "" {
		c.Next()
	} else {
		//if there is authorization header then validate it
		usr, err := validateToken(token)
		if err != nil {
			utils.ResponseError(c, http.StatusUnauthorized, "token invalid")
			c.Abort()
		}
		//if valid set context UserInfo is user object (but with field only id and role)
		c.Set(constantname.USER, usr)
		c.Next()
	}
}

func validateToken(tokenString string) (models.User, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(SECRET), nil
	})
	if err != nil {
		return models.User{}, err
	}
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok || !token.Valid {
		return models.User{}, errors.New("token invalid")
	}

	return models.User{
		ID:   uint(claims["id"].(float64)),
		Role: claims["role"].(string),
	}, nil
}
