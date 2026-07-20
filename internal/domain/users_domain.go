package domain

import (
	"time"
)

type User struct { 
  Id int64
  Email string 
  Password string 
  CreatedAt time.Time
}

type CreateUserRequest struct { 
  Id int64 `json:"id"`
  Email string `json:"email"`
  Password string `json:"password"`
}

type LoginRequest struct { 
  Email string `json:"email"`
  Password string `json:"password"`
}


// type Repository interface { 
//   Create(req *CreateUserRequest) (*domain.User, error) 
//   FindByEmail(req FindEmail) (*domain.User, error)
// }