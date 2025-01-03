package main

import (
	"context"
	"github.com/aarrico/pocket-monster-api/internal/db"
	"github.com/aarrico/pocket-monster-api/internal/utils"
	"github.com/joho/godotenv"
	"log"
	"path/filepath"
)

func main() {
	if err := godotenv.Load(filepath.Join("./", ".env")); err != nil {
		log.Fatalf("could not load .env file!")
	}

	ctx := context.Background()
	pool := utils.ConnectToDb(ctx)
	queries := db.New(pool)

	pkm, _ := queries.ListPokemon(ctx)
	print(pkm)

	pool.Close()
}
