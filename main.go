package main

import (
	"log"
	"os"

	"github.com/andil-id/api/config"
	"github.com/andil-id/api/controller"
	"github.com/andil-id/api/repository"
	"github.com/andil-id/api/router"
	"github.com/andil-id/api/service"
	"github.com/cloudinary/cloudinary-go/v2"
	"github.com/go-playground/validator/v10"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	db := config.Connection()
	validate := validator.New()
	cld, err := cloudinary.NewFromParams(config.CloudinaryCloudName, config.CloudinaryApiKey, config.CloudinaryApiSecreet)
	if err != nil {
		log.Fatalf("Error when create cloudinary instance, %v", err)
	}

	userRepository := repository.NewUserRepository()
	adminRepository := repository.NewAdminRepository()
	eventRepository := repository.NewEventRepository()

	userService := service.NewUserService(userRepository, db, validate)
	adminService := service.NewAdminService(adminRepository, db, validate, userRepository)
	authService := service.NewAuthService(db, validate, userRepository, adminRepository)
	eventService := service.NewEventService(db, validate, eventRepository, cld)

	adminController := controller.NewAdminController(adminService)
	userController := controller.NewUserController(userService)
	authController := controller.NewAuthController(authService, userService, adminService)
	eventController := controller.NewEventController(eventService)

	router := router.NewRouter(userController, adminController, authController, eventController)
	router.Run(":" + os.Getenv("PORT"))
}
