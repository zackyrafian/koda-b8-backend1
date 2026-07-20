package service

import (
	"koda-b8-backend1/internal/domain"
	"koda-b8-backend1/internal/repository"
)

type UserService struct { 
  repository *repository.UserRepository
}

func NewUserService(repository *repository.UserRepository) *UserService { 
  return &UserService{repository: repository}
}

func (s *UserService) Create(data *domain.CreateUserRequest) (*domain.User, error) {
    return s.repository.Create(data)
}

func (s *UserService) GetUsers() (*[]domain.User, error) { 
  return s.repository.FindAll()
}