package repository

import (
	"koda-b8-backend1/internal/domain"
)


type UserRepository struct {
    data *[]domain.User
}

func NewUserRepository(data *[]domain.User) *UserRepository {
    return &UserRepository{data: data}
}

func (r *UserRepository) Create(req *domain.CreateUserRequest) (*domain.User, error) {
    user := domain.User{
        Email:    req.Email,
        Password: req.Password,
    }
    *r.data = append(*r.data, user)
    return &user, nil
}


func (r *UserRepository) FindAll() (*[]domain.User, error) { 
  users := r.data
  return users, nil
}
