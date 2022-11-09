package pkg

import (
	"os"
	"time"

	"github.com/andil-id/api/model/web"
	"github.com/golang-jwt/jwt/v4"
)

func GenereateJwtToken(id string, name string, email string, role string) (string, error) {
	claims := web.Claims{
		Id:   id,
		Name: name,
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:    "BimantaraDev",
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 168)),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	jwt_secret := os.Getenv("JWT_SECRET")
	signedToken, err := token.SignedString([]byte(jwt_secret))
	if err != nil {
		return "", err
	}
	return signedToken, nil
}
