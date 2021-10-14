package middleware

import (
	constantname "bloggo/constantName"
	"bloggo/utils"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
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
		usr, err := utils.ValidateToken(token)
		if err != nil {
			utils.ResponseError(c, http.StatusUnauthorized, "token invalid")
			c.Abort()
		}
		//if valid set context UserInfo is user object (but with field only id and role)
		c.Set(constantname.USER_CTX, usr)
		c.Next()
	}
}
