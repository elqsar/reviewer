package core

import (
	"os"
	"reviewer/entity"
	"strconv"
	"time"

	"github.com/dgrijalva/jwt-go"
)

func CreateToken(user *entity.User) (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["username"] = user.Username
	claims["exp"] = time.Now().Add(1 * time.Hour).Unix()
	claims["sub"] = strconv.Itoa(user.Id)
	claims["email"] = user.Email
	secret := os.Getenv("SECRET")
	return token.SignedString([]byte(secret))
}
