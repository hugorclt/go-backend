package service

import "go-backend/src/repository"

type AuthService struct {
	Repository *repository.AuthRepository
}

func NewAuthService(repo *repository.AuthRepository) AuthService {
	return AuthService{Repository: repo}
}
