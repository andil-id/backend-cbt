package router

import (
	"fmt"
	"io"
	"os"
	"time"

	"github.com/andil-id/api/controller"
	"github.com/andil-id/api/exception"
	"github.com/andil-id/api/middleware"
	"github.com/gin-gonic/gin"
)

func NewRouter(penggunaController controller.UserController, pengurusController controller.AdminController, authController controller.AuthController, eventController controller.EventController) *gin.Engine {
	gin.SetMode(os.Getenv("GIN_MODE"))
	f, _ := os.Create(os.Getenv("PATH_LOG"))
	gin.DefaultWriter = io.MultiWriter(f, os.Stdout)
	router := gin.New()
	router.Use(middleware.CORSMiddleware())
	router.Use(exception.ErrorAppHandler())
	router.Use(gin.LoggerWithFormatter(func(param gin.LogFormatterParams) string {
		return fmt.Sprintf("%s - [%s] \"%s %s %s %d %s \"%s\" %s\"\n",
			param.ClientIP,
			param.TimeStamp.Format(time.RFC1123),
			param.Method,
			param.Path,
			param.Request.Proto,
			param.StatusCode,
			param.Latency,
			param.Request.UserAgent(),
			param.ErrorMessage,
		)
	}))
	router.Use(gin.Recovery())
	router.GET("/", func(c *gin.Context) {
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
