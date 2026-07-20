package service

import (
	"errors"
	"koda-b8-backend1/internal/domain"
	"koda-b8-backend1/internal/repository"
	"koda-b8-backend1/internal/utils"
)

type UserService struct { 
  repository *repository.UserRepository
}

func NewUserService(repository *repository.UserRepository) *UserService { 
  return &UserService{repository: repository}
}

func (s *UserService) Create(data *domain.CreateUserRequest) (*domain.User, error) {
    if !utils.IsEmailValid(data.Email) { 
      return &domain.User{}, errors.New("Email tidak sesuai")
    }
    if len(data.Password) < 8 { 
      return &domain.User{}, errors.New("Password Harus miniimal 8 Karakter")
    }
    return s.repository.Create(data)
}

func (s *UserService) GetUsers() (*[]domain.User, error) { 
  return s.repository.FindAll()
}

func (s *UserService) Login(data *domain.LoginRequest) (*domain.User, error) {
	if !utils.IsEmailValid(data.Email) {
		return nil, errors.New("format email tidak valid")
	}
	user, err := s.repository.FindByEmail(data.Email)
	if err != nil {
		return nil, err
	}
	if user.Password != data.Password {
		return nil, errors.New("password salah")
	}
	return user, nil
}