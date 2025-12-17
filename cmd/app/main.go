package main

import (
	"context"
	"log"

	"github.com/VladKovDev/promo-bot/internal/app"
	"github.com/joho/godotenv"
)

func main() {
	ctx := context.Background()

	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found")
	}

	if err := app.Run(ctx); err != nil {
		log.Fatal(err)
	}
}
