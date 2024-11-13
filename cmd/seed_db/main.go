package main

import (
	"context"
	"github.com/aarrico/pocket-monster-api/cmd/seed_db/pokeapi"
	"github.com/aarrico/pocket-monster-api/internal/db"
	"github.com/aarrico/pocket-monster-api/internal/utils"
	"github.com/joho/godotenv"
	"log"
	"os"
	"path/filepath"
)

func main() {
	if err := godotenv.Load(filepath.Join("./", ".env")); err != nil {
		log.Fatalf("could not load .env file!")
	}

	ctx := context.Background()
	pool := utils.ConnectToDb(ctx)
	queries := db.New(pool)

	populateCachesFromDb := false
	if len(os.Args) > 1 {
		switch os.Args[1] {
		case "true":
			populateCachesFromDb = true
		default:
			populateCachesFromDb = false
		}
	}

	pokeapi.Seed(ctx, queries, populateCachesFromDb)

	pool.Close()
}
