package repository

import (
	"context"
	"koda-b8-backend1/internal/domain"

	"github.com/jackc/pgx/v5/pgxpool"
)


type UserRepository struct {
    db *pgxpool.Pool
}

func NewUserRepository(db *pgxpool.Pool) *UserRepository {
    return &UserRepository{db: db}
}

func (r *UserRepository) Create(req *domain.CreateUserRequest, ctx context.Context) (*domain.User, error) {
    var users domain.User
    err := r.db.QueryRow(
      ctx, `
        INSERT INTO users (email, password) VALUES ($1, $2) RETURNING users.email, id
      `, req.Email, req.Password, 
    ).Scan(&users.Email, &users.Id)
    if err != nil { 
      return nil, err
    }
    return &users, nil
}

func (r *UserRepository) FindAll(ctx context.Context) (*[]domain.User, error) { 
  return nil, nil
}

// func (r *UserRepository) FindByEmail(email string) (*domain.User, error) {
// 	for _, user := range *r.data {
// 		if user.Email == email {
// 			return &user, nil
// 		}
// 	}

// 	return nil, errors.New("user tidak ditemukan")
// }