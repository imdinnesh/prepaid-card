package services

import (
	"github.com/imdinnesh/prepaid-card/services/auth/internal/repository"
	models "github.com/imdinnesh/prepaid-card/services/auth/model"
)

type UserService interface{
	RegisterUser(phone string) error

}

type userService struct{
	repo repository.UserRepository
}

func NewUserService(repo repository.UserRepository) UserService{
	return &userService{
		repo:repo,
	}
}

func (s *userService) RegisterUser(phone string) error{
	user:=&models.User{
		Phone: &phone,
	}
	return s.repo.CreateUser(user)
}

