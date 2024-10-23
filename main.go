package main

import (
	"context"
	"fmt"
	"github.com/aarrico/pocket-monster-api/internal/db"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/joho/godotenv"
	"log"
	"os"
)

func main() {
	if err := godotenv.Load(".env"); err != nil {
		log.Fatalf("could not load .env file!")
	}

	ctx := context.Background()
	dbConnStr := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=%s",
		os.Getenv("DB_USER"), os.Getenv("DB_PASS"), os.Getenv("DB_HOST"), os.Getenv("DB_PORT"),
		os.Getenv("DB_NAME"), os.Getenv("DB_SSL_MODE"))

	pool, err := pgxpool.New(ctx, dbConnStr)
	if err != nil {
		log.Fatal(err)
	}
	defer pool.Close()

	queries := db.New(pool)

	pkm, err := queries.ListPokemon(ctx)
	print(pkm)
}
