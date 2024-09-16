package controller

import (
	"go-backend/src/service"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
)

type UserController struct {
	Service *service.UserService
}

func NewUserController(service *service.UserService) UserController {
	return UserController{Service: service}
}

func (ctrl *UserController) InitRoute(router *chi.Mux) {
	router.Route("/users", func(router chi.Router) {
		router.Route("/{userId}", func(router chi.Router) {
			router.Get("/", ctrl.GetUserById)
		})
	})

}

func (ctrl *UserController) GetUserById(w http.ResponseWriter, r *http.Request) {
	userId := chi.URLParam(r, "userId")
	user, err := ctrl.Service.GetUserById(userId)
	if err != nil {
		http.Error(w, http.StatusText(403), 403)
		return
	}
	render.JSON(w, r, user)
}
