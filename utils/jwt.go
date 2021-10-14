package utils

import (
	constantname "bloggo/constantName"
	"bloggo/models"
	"errors"
	"fmt"
	"os"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/joho/godotenv"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

func init() {
	godotenv.Load()
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
	claims["exp"] = time.Now().Add(constantname.TOKEN_EXP_IN_HOUR * time.Hour).Unix()
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return token.SignedString([]byte(SECRET))
}

func ValidateToken(tokenString string) (models.User, error) {
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
