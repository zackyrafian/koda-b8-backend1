package main

import (
	"koda-b8-backend1/internal/di"
	"koda-b8-backend1/internal/domain"

	"github.com/gin-gonic/gin"
)

func main () {
  r := gin.Default()
  var data []domain.User
  di.Register(r, &data)
  r.Run()
}