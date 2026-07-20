package handler

import (
	"koda-b8-backend1/internal/domain"
	"koda-b8-backend1/internal/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

type UserHandler struct { 
  service *service.UserService
}

func NewUserHandler(service *service.UserService) *UserHandler { 
  return &UserHandler{service: service}
}

func (h *UserHandler) Create(c *gin.Context) { 
  email := c.PostForm("email")
  password := c.PostForm("password")
  user, err := h.service.Create(&domain.CreateUserRequest{
      Email:    email,
      Password: password,
  })
  if err != nil {
      c.JSON(http.StatusBadRequest, gin.H{
          "message": err.Error(),
      })
      return
  }
  c.JSON(http.StatusCreated, user)
}

func (h *UserHandler) Login(c *gin.Context) { 
  email := c.PostForm("email")
  password := c.PostForm("password")
  user, err := h.service.Login(&domain.LoginRequest{
    Email: email,
    Password: password,
  })

  if err != nil { 
    c.JSON(http.StatusBadRequest, gin.H{ 
      "message": err.Error(),
    })
    return
  }
  c.JSON(http.StatusAccepted, user)
}

func (h *UserHandler) GetUsers(c *gin.Context) { 
  users, err := h.service.GetUsers()
  if err != nil { 
    c.JSON(http.StatusInternalServerError, gin.H{ 
      "message": "error",
    })
    return 
  }
  c.JSON(http.StatusOK, users)
}