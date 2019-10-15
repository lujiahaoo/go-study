package utils

import (
	"time"

	"github.com/dgrijalva/jwt-go"
)

func GenerateToken(aud string) (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)
	claims := make(jwt.MapClaims)
	claims["aud"] = aud
	claims["exp"] = time.Now().Add(time.Hour.Truncate(72))
	claims["iat"] = time.Now().Unix()
	token.Claims = claims

	tokenString, err := token.SignedString([]byte(ENV("appkey").(string)))
	if err != nil {
		return "", err
	}
	return tokenString, nil
}
