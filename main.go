package main

import (
	"os"

	"github.com/andil-id/api/config"
	"github.com/andil-id/api/controller"
	"github.com/andil-id/api/repository"
	"github.com/andil-id/api/router"
	"github.com/andil-id/api/service"
	"github.com/go-playground/validator/v10"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	db := config.Connection()
	validate := validator.New()
	userRepository := repository.NewUserRepository()
	userService := service.NewUserService(userRepository, db, validate)
	userController := controller.NewUserController(userService)
	adminRepository := repository.NewAdminRepository()
	adminService := service.NewAdminService(adminRepository, db, validate, userRepository)
	adminController := controller.NewAdminController(adminService)
	authService := service.NewAuthService(db, validate, userRepository, adminRepository)
	authController := controller.NewAuthController(authService, userService, adminService)
	router := router.NewRouter(userController, adminController, authController)
	router.Run(":" + os.Getenv("PORT"))
}
