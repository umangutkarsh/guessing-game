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

type CityData struct {
	City    string   `json:"city"`
	Country string   `json:"country"`
	Clues   []string `json:"clues"`
	FunFact []string `json:"fun_fact"`
	Trivia  []string `json:"trivia"`
}

func SeedCities() error {
	db := postgresdb.GetGormClient()

	var count int64
	db.Model(&models.Destination{}).Count(&count)
	if count > 0 {
		log.Println("Cities data already seeded")
		return nil
	}

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

	for _, cityData := range cities {
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
