package controller

import (
	"github.com/gin-gonic/gin"
)

type UserController interface {
	GetUserProfile(*gin.Context)
	GetAllUserController(c *gin.Context)
	UpdateProfileUserController(c *gin.Context)
	DeleteUserController(c *gin.Context)
}
