package utils

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v5/pgxpool"
	"log"
	"os"
)

func ConnectToDb(ctx context.Context) *pgxpool.Pool {
	dbConnStr := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=%s",
		os.Getenv("DB_USER"), os.Getenv("DB_PASS"), os.Getenv("DB_HOST"), os.Getenv("DB_PORT"),
		os.Getenv("DB_NAME"), os.Getenv("DB_SSL_MODE"))

	pool, err := pgxpool.New(ctx, dbConnStr)
	if err != nil {
		log.Fatal(err)
	}

	return pool
}
