package di

import (
	"koda-b8-backend1/internal/domain"
	"koda-b8-backend1/internal/handler"
	"koda-b8-backend1/internal/repository"
	"koda-b8-backend1/internal/service"

	"github.com/gin-gonic/gin"
)

func Register(r *gin.Engine, data *[]domain.User){ 
  userRepo := repository.NewUserRepository(data)
  userService := service.NewUserService(userRepo)
  userHandler := handler.NewUserHandler(userService)

  r.POST("/users", userHandler.Create)
  r.GET("/users", userHandler.GetUsers)
}