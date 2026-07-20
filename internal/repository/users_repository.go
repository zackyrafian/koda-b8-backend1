package repository

import (
	"errors"
	"koda-b8-backend1/internal/domain"
)


type UserRepository struct {
    data *[]domain.User
}

func NewUserRepository(data *[]domain.User) *UserRepository {
    return &UserRepository{data: data}
}

func (r *UserRepository) Create(req *domain.CreateUserRequest) (*domain.User, error) {
    id := int64(len(*r.data) + 1)
    user := domain.User{
        Id: id,
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

func (r *UserRepository) FindByEmail(email string) (*domain.User, error) {
	for _, user := range *r.data {
		if user.Email == email {
			return &user, nil
		}
	}

	return nil, errors.New("user tidak ditemukan")
}