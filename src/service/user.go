package service

import (
	"go-backend/src/models"
	"go-backend/src/repository"
)

type UserService struct {
	repository *repository.UserRepository
}

func NewUserService(repo *repository.UserRepository) UserService {
	return UserService{repository: repo}
}

func (service *UserService) GetUserById(userId string) (*models.User, error) {
	return service.repository.GetUserById(userId)
}
