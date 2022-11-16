package controller

import (
	"net/http"

	"github.com/andil-id/api/exception"
	"github.com/andil-id/api/helper"
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
	var res web.LoginResponse
	user := c.Query("user")
	switch user {
	case "user":
		req := web.LoginUserRequest{}
		err := c.ShouldBindJSON(&req)
		if err != nil {
			c.Error(err)
			return
		}
		token, err := authController.AuthService.LoginUser(c.Request.Context(), req)
		if err != nil {
			c.Error(err)
			return
		}
		res.Token = token
	case "admin":
		req := web.LoginAdminRequest{}
		err := c.ShouldBindJSON(&req)
		if err != nil {
			c.Error(err)
			return
		}
		token, err := authController.AuthService.LoginAdmin(c.Request.Context(), req)
		if err != nil {
			c.Error(err)
			return
		}
		res.Token = token
	default:
		c.Error(e.Wrap(exception.ErrBadRequest, "request params not allowed!"))
		return
	}
	helper.ResponseSuccess(c, res, helper.Meta{
		StatusCode: http.StatusOK,
	})
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
		helper.ResponseSuccess(c, data, helper.Meta{
			StatusCode: http.StatusOK,
		})
	} else if user == "admin" {
		auth := web.RegisterAdminRequest{}
		err := c.ShouldBindJSON(&auth)
		if err != nil {
			c.Error(err)
			return
		}
		res, err := authController.AdminService.RegisterAdmin(c.Request.Context(), auth)
		if err != nil {
			c.Error(err)
			return
		}
		helper.ResponseSuccess(c, res, helper.Meta{
			StatusCode: http.StatusOK,
		})
	} else {
		c.Error(e.Wrap(exception.ErrBadRequest, "request param not allowed"))
	}
}
