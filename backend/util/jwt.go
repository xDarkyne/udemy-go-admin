package util

import (
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/xdarkyne/udemy/config"
)

func GenerateJWT(issuer string) (string, error) {
	claims := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.StandardClaims{
		Issuer:    issuer,
		ExpiresAt: time.Now().Add(time.Hour * 24).Unix(),
	})

	return claims.SignedString(config.App.AuthSecret)
}

func ParseJWT(cookie string) (string, error) {
	token, err := jwt.ParseWithClaims(cookie, &jwt.StandardClaims{}, func(token *jwt.Token) (interface{}, error) {
		return config.App.AuthSecret, nil
	})

	if err != nil || !token.Valid {
		return "", err
	}

	claims := token.Claims.(*jwt.StandardClaims)
	return claims.Issuer, nil
}
