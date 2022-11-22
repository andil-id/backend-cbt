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

func NewRouter(userController controller.UserController, adminController controller.AdminController, authController controller.AuthController, eventController controller.EventController, orderController controller.OrderController) *gin.Engine {
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
		api.GET("/users", middleware.JwtAuthMiddleware(), userController.GetAllUserController)
		api.DELETE("/users/:id", middleware.JwtAuthMiddleware(), userController.DeleteUserController)
		api.GET("/users/profile", middleware.JwtAuthMiddleware(), userController.GetUserProfile)
		api.PUT("/users/profile", middleware.JwtAuthMiddleware(), userController.UpdateProfileUserController)
		// * admin
		api.GET("/admins", middleware.JwtAuthMiddleware(), adminController.GetAllAdminController)
		api.GET("/admins/:id", middleware.JwtAuthMiddleware(), adminController.GetAdminByIdController)
		api.DELETE("/admins/:id", middleware.JwtAuthMiddleware(), adminController.DeleteAdminController)
		api.PUT("/admins/profile", middleware.JwtAuthMiddleware(), adminController.UpdateProfileAdminController)
		// * event
		api.POST("/events", middleware.JwtAuthMiddleware(), eventController.AddEvent)
		api.GET("/events", eventController.GetAllEvents)
		api.GET("/events/:id", eventController.GetEventById)
		events := api.Group("/events")
		{
			events.POST("/order", middleware.JwtAuthMiddleware(), orderController.CreateOrderEvent)
			events.PUT("/order/confirm/:id", middleware.JwtAuthMiddleware(), orderController.ConfirmOrder)
			events.PUT("/order/reject/:id", middleware.JwtAuthMiddleware(), orderController.RejectOrder)
		}
	}
	return router
}
