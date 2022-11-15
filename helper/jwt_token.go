package helper

import (
	"time"

	"github.com/andil-id/api/config"
	"github.com/andil-id/api/model/web"
	"github.com/golang-jwt/jwt/v4"
)

func GenereateJwtToken(id string, name string, role string) (string, error) {
	claims := web.Claims{
		Id:   id,
		Name: name,
		Role: role,
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:    "BimantaraDev",
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 168)),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	jwt_secret := config.JwtSecreet()
	signedToken, err := token.SignedString([]byte(jwt_secret))
	if err != nil {
		return "", err
	}
	return signedToken, nil
}
