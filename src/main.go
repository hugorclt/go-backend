package main

import (
	"go-backend/src/config"
	"go-backend/src/controller"
	"go-backend/src/db"
	"go-backend/src/httpServer"
	"go-backend/src/repository"
	"go-backend/src/service"
)

func main() {
	config := config.InitConfig()
	db := db.NewDB(config)
	server := httpServer.InitServer(config, db)

	userRepository := repository.NewUserRepository(db)
	authRepository := repository.NewAuthRepository(db)

	userService := service.NewUserService(&userRepository)
	authService := service.NewAuthService(&authRepository)

	userController := controller.NewUserController(&userService)
	authController := controller.NewAuthController(&authService, &userService)

	authController.InitRoute(server.Router)
	userController.InitRoute(server.Router)

	server.Run()
}
