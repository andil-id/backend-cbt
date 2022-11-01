package controller

import "github.com/gin-gonic/gin"

type AdminController interface {
	GetAdminByIdController(c *gin.Context)
	GetAllAdminController(c *gin.Context)
	UpdateProfileAdminController(c *gin.Context)
	DeleteAdminController(c *gin.Context)
}
