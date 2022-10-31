package web

import "github.com/golang-jwt/jwt/v4"

type Claims struct {
	Id   string `json:"id"`
	Nama string `json:"nama"`
	jwt.RegisteredClaims
}
