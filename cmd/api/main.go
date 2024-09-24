package main

import (
	"github.com/joho/godotenv"
	"log"
)

func main() {

	if err := godotenv.Load(); err != nil {
		log.Fatalf("could not load .env file!")
	}

}
