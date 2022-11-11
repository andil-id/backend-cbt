package controller

import (
	"net/http"

	"github.com/andil-id/api/exception"
	"github.com/andil-id/api/model/web"
	"github.com/andil-id/api/service"
	"github.com/gin-gonic/gin"
	e "github.com/pkg/errors"
)

type AuthControllerImpl struct {
	UserService  service.UserService
	AdminService service.AdminService
	AuthService  service.AuthService
}

func NewAuthController(authService service.AuthService, penggunaService service.UserService, pengurusService service.AdminService) AuthController {
	return &AuthControllerImpl{
		UserService:  penggunaService,
		AdminService: pengurusService,
		AuthService:  authService,
	}
}

func (authController AuthControllerImpl) LoginController(c *gin.Context) {
	user := c.Query("user")
	auth := web.LoginAuthRequest{}
	err := c.ShouldBindJSON(&auth)
	if err != nil {
		c.Error(err)
		return
	}
	token, err := authController.AuthService.Login(c.Request.Context(), user, auth)
	if err != nil {
		c.Error(err)
		return
	} else {
		c.JSON(http.StatusOK, gin.H{
			"code":    http.StatusOK,
			"status":  "succes",
			"message": "Login succesfully",
			"data":    token,
		})
	}
}
func (authController AuthControllerImpl) RegisterController(c *gin.Context) {
	user := c.Query("user")
	if user == "user" {
		auth := web.RegisterUserRequest{}
		err := c.ShouldBindJSON(&auth)
		if err != nil {
			c.Error(err)
			return
		}

		data, err := authController.UserService.RegisterUser(c.Request.Context(), auth)
		if err != nil {
			c.Error(err)
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"code":    http.StatusOK,
			"status":  "OK",
			"message": "Register user succesfully",
			"data":    data,
		})
	} else if user == "admin" {
		auth := web.RegisterAdminRequest{}
		err := c.ShouldBindJSON(&auth)
		if err != nil {
			c.Error(err)
			return
		}
		err = authController.AdminService.RegisterAdmin(c.Request.Context(), auth)
		if err != nil {
			c.Error(err)
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"code":    http.StatusOK,
			"status":  "OK",
			"message": "Register admin succesfully",
		})
	} else {
		c.Error(e.Wrap(exception.ErrBadRequest, "Param not allowed"))
	}
}
