package main

import (
	"context"
	"github.com/aarrico/pocket-monster-api/db/seed"
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

	//seedCmd := flag.NewFlagSet("seed", flag.ExitOnError)

	if len(os.Args) > 1 {
		switch os.Args[1] {
		case "seed":
			seed.SeedPokemon()
			os.Exit(0)
		default:
			log.Fatalf("invalid command line argument")
		}
	}

	ctx := context.Background()
	pool := utils.ConnectToDb(ctx)
	queries := db.New(pool)
	pkm, _ := queries.ListPokemon(ctx)
	print(pkm)

	pool.Close()
}
