package controller

import (
	"fmt"
	"go-backend/src/dto"
	"go-backend/src/middleware"
	"go-backend/src/service"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
)

type AuthController struct {
	AuthService *service.AuthService
	UserService *service.UserService
}

func NewAuthController(authService *service.AuthService, userService *service.UserService) AuthController {
	return AuthController{AuthService: authService, UserService: userService}
}

func (controller *AuthController) InitRoute(router *chi.Mux) {
	router.Route("/auth", func(router chi.Router) {
		router.Post("/login", controller.LoginHandler)
		router.With(middleware.ValidateDTO(&dto.RegisterDTO{})).Post("/register", controller.RegisterHandler)
	})
}

func (controller *AuthController) LoginHandler(res http.ResponseWriter, req *http.Request) {
}

func (controller *AuthController) RegisterHandler(res http.ResponseWriter, req *http.Request) {
	registerDto := middleware.GetDTOFromContext(req)

	controller.UserService.CreateUser(registerDto)
	render.PlainText(res, req, "OK")
}
