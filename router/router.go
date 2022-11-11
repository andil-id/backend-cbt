package router

import (
	"io"
	"log"
	"os"

	"github.com/andil-id/api/config"
	"github.com/andil-id/api/controller"
	"github.com/andil-id/api/middleware"
	"github.com/gin-gonic/gin"
)

func NewRouter(penggunaController controller.UserController, pengurusController controller.AdminController, authController controller.AuthController, eventController controller.EventController) *gin.Engine {
	gin.SetMode(config.GinMode)
	f, err := os.Create(config.PathLog)
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
		log.Println(c.Request.Method)
		c.JSON(200, gin.H{
			"message": "ok",
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
		api.GET("/user", middleware.JwtAuthMiddleware(), penggunaController.GetUserByIdController)
		api.DELETE("/user", middleware.JwtAuthMiddleware(), penggunaController.DeleteUserController)
		api.PUT("/user/profile", middleware.JwtAuthMiddleware(), penggunaController.UpdateProfileUserController)
		api.GET("/all/user", middleware.JwtAuthMiddleware(), penggunaController.GetAllUserController)
		// * admin
		api.GET("/admin", middleware.JwtAuthMiddleware(), pengurusController.GetAdminByIdController)
		api.DELETE("/admin", middleware.JwtAuthMiddleware(), pengurusController.DeleteAdminController)
		api.PUT("/admin/profile", middleware.JwtAuthMiddleware(), pengurusController.UpdateProfileAdminController)
		api.GET("/all/admin", middleware.JwtAuthMiddleware(), pengurusController.GetAllAdminController)
		// * event
		api.POST("/event/add", eventController.AddEvent)
	}
	return router
}
