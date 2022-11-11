package middleware

import (
	"fmt"
	"net/http"
	"os"
	"strings"

	"github.com/andil-id/api/helper"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
)

func JwtAuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		authorizationToken := c.GetHeader("Authorization")
		// * return invalid if bearer type not set in token
		if !strings.Contains(authorizationToken, "Bearer") {
			helper.ResponseError(c, http.StatusUnauthorized, "Authorization type not supported")
			return
		}
		tokenString := strings.Replace(authorizationToken, "Bearer ", "", -1)

		// * check signing method in token
		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (any, error) {
			if method, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("signing method invalid")
			} else if method != jwt.SigningMethodHS256 {
				return nil, fmt.Errorf("signing method invalid")
			}
			jwt_secret := os.Getenv("JWT_SECRET")
			return []byte(jwt_secret), nil
		})
		if err != nil {
			helper.ResponseError(c, http.StatusUnauthorized, "Token is invalid")
			return
		}

		// * validation token
		claims, ok := token.Claims.(jwt.MapClaims)
		if !ok || !token.Valid {
			helper.ResponseError(c, http.StatusUnauthorized, "Token is invalid")
			return
		}
		// * set token to context
		c.Set("token", claims)
		c.Next()
	}
}
