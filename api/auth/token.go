package auth

import (
	"github.com/dgrijalva/jwt-go"
	"os"
	"time"
)

func CreateToken(userid uint32) (string, error) {
	claims := jwt.MapClaims{}
	claims["authorized"] 	= true
	claims["user_id"]		= userid
	claims["exp"]			= time.Now().Add(time.Hour * 1).Unix()

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, err := token.SignedString([]byte(os.Getenv("API_SECRET")))
	return tokenString, err
}
