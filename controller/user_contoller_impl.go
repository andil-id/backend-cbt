package controller

import (
	"fmt"
	"net/http"

	"github.com/andil-id/api/exception"
	"github.com/andil-id/api/helper"
	"github.com/andil-id/api/model/web"
	"github.com/andil-id/api/service"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	e "github.com/pkg/errors"
)

type UserControllerImpl struct {
	UserService service.UserService
}

func NewUserController(userService service.UserService) UserController {
	return &UserControllerImpl{
		UserService: userService,
	}
}

func (userContoller UserControllerImpl) GetUserByIdController(c *gin.Context) {
	token := c.MustGet("token").(jwt.MapClaims)
	id := token["id"]

	res, err := userContoller.UserService.GetUserById(c.Request.Context(), fmt.Sprintf("%v", id))
	if err != nil {
		c.Error(err)
		return
	}
	helper.ResponseSuccess(c, res, helper.Meta{
		StatusCode: http.StatusOK,
	})
}

func (userContoller UserControllerImpl) GetAllUserController(c *gin.Context) {
	token := c.MustGet("token").(jwt.MapClaims)
	if token["role"].(string) != "admin" {
		c.Error(e.Wrap(exception.ErrUnAuth, ""))
		return
	}

	res, err := userContoller.UserService.GetAllUser(c.Request.Context())
	if err != nil {
		c.Error(err)
		return
	}
	helper.ResponseSuccess(c, res, helper.Meta{
		StatusCode: http.StatusOK,
		TotalData:  len(res),
	})
}

func (userContoller UserControllerImpl) DeleteUserController(c *gin.Context) {
	token := c.MustGet("token").(jwt.MapClaims)
	id := token["id"]

	err := userContoller.UserService.DeleteUser(c.Request.Context(), fmt.Sprintf("%v", id))
	if err != nil {
		c.Error(err)
		return
	}
	helper.ResponseSuccess(c, nil, helper.Meta{
		StatusCode: http.StatusOK,
	})
}

func (userContoller UserControllerImpl) UpdateProfileUserController(c *gin.Context) {
	token := c.MustGet("token").(jwt.MapClaims)
	id := token["id"]

	req := web.UpdateProfileUserRequest{}
	err := c.ShouldBindJSON(&req)
	if err != nil {
		c.Error(err)
		return
	}
	err = userContoller.UserService.UpdateProfileUser(c.Request.Context(), fmt.Sprintf("%v", id), req)
	if err != nil {
		c.Error(err)
		return
	}
	helper.ResponseSuccess(c, nil, helper.Meta{
		StatusCode: http.StatusOK,
	})
}
