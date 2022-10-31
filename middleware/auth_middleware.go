package middleware

import (
	"fmt"
	"net/http"
	"os"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
)

func JwtAuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		authorizationToken := c.GetHeader("Authorization")
		// * return invalid if bearer type not set in token
		if !strings.Contains(authorizationToken, "Bearer") {
			c.JSON(http.StatusUnauthorized, gin.H{
				"code":    http.StatusUnauthorized,
				"status":  "failed",
				"message": "Invalid token",
				"data":    nil,
			})
			c.Abort()
			return
		}
		tokenString := strings.Replace(authorizationToken, "Bearer ", "", -1)

		// * check signing method in token
		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			if method, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("signing method invalid")
			} else if method != jwt.SigningMethodHS256 {
				return nil, fmt.Errorf("signing method invalid")
			}
			jwt_secret := os.Getenv("JWT_SECRET")
			return []byte(jwt_secret), nil
		})

		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{
				"code":    http.StatusUnauthorized,
				"status":  "failed",
				"message": err.Error(),
				"data":    nil,
			})
			c.Abort()
			return
		}

		// * validation token
		claims, ok := token.Claims.(jwt.MapClaims)
		if !ok || !token.Valid {
			c.JSON(http.StatusUnauthorized, gin.H{
				"code":    http.StatusUnauthorized,
				"status":  "failed",
				"message": "Invalid Token",
				"data":    nil,
			})
			c.Abort()
			return
		}
		// * set token to context
		c.Set("token", claims)
		c.Next()
	}
}
