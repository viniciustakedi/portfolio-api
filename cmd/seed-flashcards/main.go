package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"os"
	"portfolio/api/api/flashcards"
	"portfolio/api/config"
	"portfolio/api/infra/db"
)

func main() {
	environment := flag.String("e", "development", "Environment (development, production)")
	force := flag.Bool("force", false, "Delete existing flashcard collections and re-seed")
	flag.Parse()

	config.Init(*environment)

	if err := db.InitMongoDB(); err != nil {
		log.Fatalf("mongodb: %v", err)
	}
	defer db.KillMongoDB()

	ctx := context.Background()
	err := flashcards.SeedCollections(ctx, db.GetMongoDB(), *force)
	if err != nil {
		if err == flashcards.ErrAlreadySeeded {
			fmt.Println(err.Error())
			os.Exit(0)
		}
		log.Fatalf("seed: %v", err)
	}
	fmt.Println("Flashcard paths and cards seeded successfully.")
}
