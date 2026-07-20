package main

import (
	"context"
	"fmt"
	"koda-b8-backend1/internal/di"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgxpool"
)

func main () {
  r := gin.Default()
  conn, err := pgxpool.New(context.Background(), "postgres://postgres:admin@localhost:5433/postgres")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}
	defer conn.Close()
  di.Register(r, conn)
  r.Run(":9999")
}