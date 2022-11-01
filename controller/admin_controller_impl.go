package controller

import (
	"fmt"
	"net/http"

	"github.com/andil-id/api/model/web"
	"github.com/andil-id/api/service"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
)

type AdminControllerImpl struct {
	AdminService service.AdminService
}

func NewAdminController(adminService service.AdminService) AdminController {
	return &AdminControllerImpl{
		AdminService: adminService,
	}
}
func RegisterAdminController(adminService service.AdminService) AdminController {
	return &AdminControllerImpl{
		AdminService: adminService,
	}
}

func (adminController *AdminControllerImpl) GetAdminByIdController(c *gin.Context) {
	token := c.MustGet("token").(jwt.MapClaims)
	id := token["id"]
	if token["role"].(string) == "admin" {
		response, err := adminController.AdminService.GetAdminById(c.Request.Context(), fmt.Sprintf("%v", id))
		if err != nil {
			c.Error(err)
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"code":    http.StatusOK,
			"status":  "succes",
			"message": "Get admin by id succesfully",
			"data":    response,
		})
	} else {
		c.JSON(http.StatusUnauthorized, gin.H{
			"code":    http.StatusUnauthorized,
			"message": "Your'e Not Authorized",
		})
	}
}

func (adminController *AdminControllerImpl) GetAllAdminController(c *gin.Context) {
	token := c.MustGet("token").(jwt.MapClaims)
	if token["role"].(string) == "admin" {
		response, err := adminController.AdminService.GetAllAdmin(c.Request.Context())
		if err != nil {
			c.Error(err)
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"code":    http.StatusOK,
			"status":  "succes",
			"message": "Get all admin succesfully",
			"data":    response,
		})
	} else {
		c.JSON(http.StatusUnauthorized, gin.H{
			"code":    http.StatusUnauthorized,
			"message": "Your'e Not Authorized",
		})
	}
}
func (adminController *AdminControllerImpl) UpdateProfileAdminController(c *gin.Context) {
	token := c.MustGet("token").(jwt.MapClaims)
	id := token["id"]
	if token["role"].(string) == "admin" {
		admin := web.UpdateProfileAdminRequest{}
		err := c.ShouldBindJSON(&admin)
		if err != nil {
			c.Error(err)
			return
		}
		err = adminController.AdminService.UpdateProfileAdmin(c.Request.Context(), fmt.Sprintf("%v", id), admin)
		if err != nil {
			c.Error(err)
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"code":    http.StatusOK,
			"status":  "succes",
			"message": "Update profile admin succesfully",
		})
	} else {
		c.JSON(http.StatusUnauthorized, gin.H{
			"code":    http.StatusUnauthorized,
			"message": "Your'e Not Authorized",
		})
	}
}
func (adminController *AdminControllerImpl) DeleteAdminController(c *gin.Context) {
	token := c.MustGet("token").(jwt.MapClaims)
	id := token["id"]
	if token["role"].(string) == "admin" {
		adminController.AdminService.DeleteAdmin(c.Request.Context(), fmt.Sprintf("%v", id))
		c.JSON(200, gin.H{
			"code":   200,
			"status": "Delete user succesfully",
		})
	} else {
		c.JSON(http.StatusUnauthorized, gin.H{
			"code":    http.StatusUnauthorized,
			"message": "Your'e Not Authorized",
		})
	}
}
