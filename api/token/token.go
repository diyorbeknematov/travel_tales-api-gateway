package token

import (
	"api-gateway/config"
	"errors"

	"github.com/dgrijalva/jwt-go"
)

type Claims struct {
	UserId   string
	Username string
	Email    string
	jwt.StandardClaims
}

func ExtractClaimsAccess(tokenString string) (*Claims, error) {
	cfg := config.Load()
	claims := &Claims{}

	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(cfg.ACCESS_TOKEN), nil
	})

	if err != nil {
		return nil, err
	}

	if !token.Valid {
		return nil, errors.New("invalid token")
	}

	return claims, nil
}

func ValidateToken(tokenStr string) (bool, error) {
	_, err := ExtractClaimsAccess(tokenStr)
	if err != nil {
		return false, err
	}
	return true, nil
}
