package services

import (
	"time"

	"guessing-game/internal/domain/models"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type ChallengeService struct {
	DB *gorm.DB
}

func NewChallengeService(db *gorm.DB) *ChallengeService {
	return &ChallengeService{DB: db}
}

func (s *ChallengeService) CreateChallenge(userID uuid.UUID) (*models.Challenge, error) {
	token := uuid.New().String()
	challenge := models.Challenge{
		ID:            uuid.New(),
		InviterUserID: userID,
		Token:         token,
		InviteLink:    "http://whatsapp.com/challenge/" + token,
		CreatedAt:     time.Now(),
	}
	if err := s.DB.Create(&challenge).Error; err != nil {
		return nil, err
	}
	return &challenge, nil
}

func (s *ChallengeService) GetChallenge(token string) (*models.Challenge, error) {
	var challenge models.Challenge
	if err := s.DB.Where("token = ?", token).First(&challenge).Error; err != nil {
		return nil, err
	}
	return &challenge, nil
}
