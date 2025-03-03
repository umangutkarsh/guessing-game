// cmd/main.go
package main

import (
	"guessing-game/internal/adapter/infrastructure/postgresdb"
	"guessing-game/internal/adapter/infrastructure/postgresdb/seed"
	"guessing-game/internal/bootstrap"
	"guessing-game/internal/handlers"
	"guessing-game/internal/services"
	"log"

	"github.com/joho/godotenv"
)

func init() {
	// Load .env file if exists
	if err := godotenv.Load(); err != nil {
		log.Println("Warning: .env file not found")
	}

	postgresdb.InitGormClient()

	db := postgresdb.GetGormClient()
	handlers.AuthService = services.NewAuthService(db)
	handlers.GameService = services.NewGameService(db)
	handlers.ChallengeService = services.NewChallengeService(db)

	if err := seed.SeedCities(); err != nil {
		log.Fatalf("Failed to seed cities data: %v", err)
	}
}

func main() {
	router := bootstrap.NewRouter()

	if err := router.Run(":8080"); err != nil {
		log.Fatalf("Error starting server: %v", err)
	}
}
