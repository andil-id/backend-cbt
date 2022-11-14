package router

import (
	"io"
	"log"
	"os"

	"github.com/andil-id/api/config"
	"github.com/andil-id/api/controller"
	"github.com/andil-id/api/helper"
	"github.com/andil-id/api/middleware"
	"github.com/gin-gonic/gin"
)

func NewRouter(penggunaController controller.UserController, pengurusController controller.AdminController, authController controller.AuthController, eventController controller.EventController) *gin.Engine {
	gin.SetMode(config.GinMode())
	f, err := os.Create(config.PathLog())
	if err != nil {
		log.Fatalf("Error when initialize path log %v", err)
	}
	gin.DefaultWriter = io.MultiWriter(f, os.Stdout)

	router := gin.New()
	router.Use(gin.Recovery())
	router.Use(middleware.CORSMiddleware())
	router.Use(middleware.ErrorAppHandler())
	router.Use(gin.LoggerWithFormatter(middleware.Loogger))

	router.GET("/", func(c *gin.Context) {
		helper.ResponseSuccess(c, nil, helper.Meta{
			StatusCode: 200,
		})
	})

	api := router.Group("/api")
	{
		auth := api.Group("/auth")
		{
			auth.POST("/login", authController.LoginController)
			auth.POST("/register", authController.RegisterController)
		}
		// * user
		api.GET("/users", middleware.JwtAuthMiddleware(), penggunaController.GetAllUserController)
		api.GET("/users/:id", middleware.JwtAuthMiddleware(), penggunaController.GetUserByIdController)
		api.DELETE("/users/:id", middleware.JwtAuthMiddleware(), penggunaController.DeleteUserController)
		api.PUT("/users/profile", middleware.JwtAuthMiddleware(), penggunaController.UpdateProfileUserController)
		// * admin
		api.GET("/admins", middleware.JwtAuthMiddleware(), pengurusController.GetAllAdminController)
		api.GET("/admins/:id", middleware.JwtAuthMiddleware(), pengurusController.GetAdminByIdController)
		api.DELETE("/admins/:id", middleware.JwtAuthMiddleware(), pengurusController.DeleteAdminController)
		api.PUT("/admins/profile", middleware.JwtAuthMiddleware(), pengurusController.UpdateProfileAdminController)
		// * event
		api.POST("/events", middleware.JwtAuthMiddleware(), eventController.AddEvent)
		api.GET("/events", middleware.JwtAuthMiddleware(), eventController.GetAllEvents)
		api.GET("/events/:id", middleware.JwtAuthMiddleware(), eventController.GetEventById)
	}
	return router
}
