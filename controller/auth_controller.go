package controller

import "github.com/gin-gonic/gin"

type AuthController interface {
	LoginController(c *gin.Context)
	RegisterController(c *gin.Context)
}
