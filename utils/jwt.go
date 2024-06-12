package utils

import (
	"github.com/golang-jwt/jwt/v5"
	"time"
)

const SecretKey = "secret"

func GenerateJwt(issuer string) (string, error) {

	expirationTime := time.Now().Add(time.Hour * 24)
	claims := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.RegisteredClaims{
		Issuer:    issuer,
		ExpiresAt: jwt.NewNumericDate(expirationTime),
	})
	return claims.SignedString([]byte(SecretKey))

}

type Claims struct {
	jwt.RegisteredClaims
}

func ParseJwt(cookie string) (string, error) {
	token, err := jwt.ParseWithClaims(cookie, &jwt.RegisteredClaims{}, func(token *jwt.Token) (interface{}, error) {

		return []byte(SecretKey), nil
	})

	if err != nil || !token.Valid {
		return "", err
	}

	claims := token.Claims.(*jwt.RegisteredClaims)
	return claims.Issuer, nil
}
