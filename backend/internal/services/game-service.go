package services

import (
	"encoding/json"
	"errors"
	"math/rand"
	"time"

	"guessing-game/internal/domain/models"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type GameService struct {
	DB *gorm.DB
}

func NewGameService(db *gorm.DB) *GameService {
	return &GameService{DB: db}
}

type RandomDestinationResponse struct {
	QuestionToken string   `json:"questionToken"`
	Clues         []string `json:"clues"`
	Options       []Option `json:"options"`
}

type Option struct {
	CityID   uuid.UUID `json:"cityId"`
	CityName string    `json:"cityName"`
}

func (s *GameService) GetRandomDestination() (*RandomDestinationResponse, error) {
	var destination models.Destination
	if err := s.DB.Order("RANDOM()").First(&destination).Error; err != nil {
		return nil, err
	}

	var clues []string
	if err := json.Unmarshal(destination.Clues, &clues); err != nil {
		return nil, err
	}

	rand.Seed(time.Now().UnixNano())
	nClues := 1
	if len(clues) > 1 && rand.Intn(2) == 1 {
		nClues = 2
	}
	selectedClues := clues[:nClues]

	// Build options: include the correct destination and three decoys.
	var options []Option
	correctOption := Option{
		CityID:   destination.ID,
		CityName: destination.City,
	}
	options = append(options, correctOption)

	var decoys []models.Destination
	if err := s.DB.Where("id <> ?", destination.ID).Order("RANDOM()").Limit(3).Find(&decoys).Error; err != nil {
		return nil, err
	}
	for _, d := range decoys {
		options = append(options, Option{
			CityID:   d.ID,
			CityName: d.City,
		})
	}

	rand.Shuffle(len(options), func(i, j int) {
		options[i], options[j] = options[j], options[i]
	})

	resp := RandomDestinationResponse{
		QuestionToken: destination.ID.String(), // use UUID string as token
		Clues:         selectedClues,
		Options:       options,
	}
	return &resp, nil
}

type GuessResult struct {
	Correct bool   `json:"correct"`
	FunFact string `json:"funFact"`
	Trivia  string `json:"trivia"`
	Score   int    `json:"score"`
}

func (s *GameService) ProcessGuess(userID uuid.UUID, questionToken string, selectedCityID uuid.UUID) (*GuessResult, error) {
	correctID, err := uuid.Parse(questionToken)
	if err != nil {
		return nil, errors.New("invalid question token")
	}

	var destination models.Destination
	if err := s.DB.First(&destination, "id = ?", correctID).Error; err != nil {
		return nil, err
	}

	var funFacts []string
	if err := json.Unmarshal(destination.FunFact, &funFacts); err != nil {
		return nil, err
	}
	var trivia []string
	if err := json.Unmarshal(destination.Trivia, &trivia); err != nil {
		return nil, err
	}

	rand.Seed(time.Now().UnixNano())
	funFact := ""
	if len(funFacts) > 0 {
		funFact = funFacts[rand.Intn(len(funFacts))]
	}
	triviaStr := ""
	if len(trivia) > 0 {
		triviaStr = trivia[rand.Intn(len(trivia))]
	}

	correct := selectedCityID == destination.ID

	var user models.User
	if err := s.DB.First(&user, "id = ?", userID).Error; err != nil {
		return nil, err
	}
	if correct {
		user.Score++
		s.DB.Save(&user)
	}

	result := GuessResult{
		Correct: correct,
		FunFact: funFact,
		Trivia:  triviaStr,
		Score:   user.Score,
	}
	return &result, nil
}

func (s *GameService) GetScore(userID uuid.UUID) (int, error) {
	var user models.User
	if err := s.DB.First(&user, "id = ?", userID).Error; err != nil {
		return 0, err
	}
	return user.Score, nil
}
