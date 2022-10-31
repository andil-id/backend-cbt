package controller

import (
	"fmt"
	"net/http"

	"github.com/andil-id/api/model/web"
	"github.com/andil-id/api/service"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
)

type UserControllerImpl struct {
	UserService service.UserService
}

func NewUserController(userService service.UserService) UserController {
	return &UserControllerImpl{
		UserService: userService,
	}
}
func RegisterUserController(userService service.UserService) UserController {
	return &UserControllerImpl{
		UserService: userService,
	}
}
func (userContoller UserControllerImpl) GetUserByIdController(c *gin.Context) {
	token := c.MustGet("token").(jwt.MapClaims)
	id := token["id"]
	if token["role"].(string) == "user" {
		user, err := userContoller.UserService.GetUserById(c.Request.Context(), fmt.Sprintf("%v", id))
		if err != nil {
			c.Error(err)
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"code":    http.StatusOK,
			"status":  "succes",
			"message": "Get user by id succesfully",
			"data":    user,
		})
	} else {
		c.JSON(http.StatusUnauthorized, gin.H{
			"code":    http.StatusUnauthorized,
			"message": "Your'e Not Authorized",
		})
	}
}
func (userContoller UserControllerImpl) GetAllUserController(c *gin.Context) {
	token := c.MustGet("token").(jwt.MapClaims)
	if token["role"].(string) == "admin" {
		user, err := userContoller.UserService.GetAllUser(c.Request.Context())
		if err != nil {
			c.Error(err)
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"code":    http.StatusOK,
			"status":  "succes",
			"message": "Get all user succesfully",
			"data":    user,
		})
	} else {
		c.JSON(http.StatusUnauthorized, gin.H{
			"code":    http.StatusUnauthorized,
			"message": "Your'e Not Authorized",
		})
	}

}
func (userContoller UserControllerImpl) DeleteUserController(c *gin.Context) {
	token := c.MustGet("token").(jwt.MapClaims)
	id := token["id"]
	if token["role"].(string) == "user" {
		userContoller.UserService.DeleteUser(c.Request.Context(), fmt.Sprintf("%v", id))
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
func (userContoller UserControllerImpl) UpdateProfileUserController(c *gin.Context) {
	token := c.MustGet("token").(jwt.MapClaims)
	id := token["id"]
	if token["role"].(string) == "user" {
		user := web.UpdateProfileUserRequest{}
		err := c.ShouldBindJSON(&user)
		if err != nil {
			c.Error(err)
			return
		}
		err = userContoller.UserService.UpdateProfileUser(c.Request.Context(), fmt.Sprintf("%v", id), user)
		if err != nil {
			c.Error(err)
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"code":    http.StatusOK,
			"status":  "succes",
			"message": "Update profile user succesfully",
			"data":    nil,
		})
	} else {
		c.JSON(http.StatusUnauthorized, gin.H{
			"code":    http.StatusUnauthorized,
			"message": "Your'e Not Authorized",
		})
	}
}
