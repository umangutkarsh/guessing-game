// internal/adapter/infrastructure/postgresdb/seed.go
package seed

import (
	"encoding/json"
	"guessing-game/internal/adapter/infrastructure/postgresdb"
	"guessing-game/internal/domain/models"
	"io/ioutil"
	"log"
	"os"

	"github.com/google/uuid"
)

// CityData represents the JSON structure for seeding.
type CityData struct {
	City    string   `json:"city"`
	Country string   `json:"country"`
	Clues   []string `json:"clues"`
	FunFact []string `json:"fun_fact"`
	Trivia  []string `json:"trivia"`
}

// SeedCities reads data/cities.json and seeds the database.
func SeedCities() error {
	// Get Gorm DB client
	db := postgresdb.GetGormClient()

	// Check if data is already seeded
	var count int64
	db.Model(&models.Destination{}).Count(&count)
	if count > 0 {
		log.Println("Cities data already seeded")
		return nil
	}

	// Read the JSON file (ensure file exists at data/cities.json)
	file, err := os.Open("data/cities.json")
	if err != nil {
		return err
	}
	defer file.Close()

	bytes, err := ioutil.ReadAll(file)
	if err != nil {
		return err
	}

	var cities []CityData
	if err := json.Unmarshal(bytes, &cities); err != nil {
		return err
	}

	// Iterate and insert each city into the database
	for _, cityData := range cities {
		// Marshal the slices into JSON for storing in the JSONB columns
		clues, _ := json.Marshal(cityData.Clues)
		funFacts, _ := json.Marshal(cityData.FunFact)
		trivia, _ := json.Marshal(cityData.Trivia)

		city := models.Destination{
			ID:      uuid.New(),
			City:    cityData.City,
			Country: cityData.Country,
			Clues:   clues,
			FunFact: funFacts,
			Trivia:  trivia,
		}

		if err := db.Create(&city).Error; err != nil {
			return err
		}
	}
	log.Println("Seeded cities data successfully")
	return nil
}
