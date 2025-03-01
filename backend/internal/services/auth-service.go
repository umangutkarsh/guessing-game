package services

import (
	"errors"
	"guessing-game/internal/domain/models"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type AuthService struct {
	DB *gorm.DB
}

func NewAuthService(db *gorm.DB) *AuthService {
	return &AuthService{DB: db}
}

func (s *AuthService) RegisterUser(username string) (*models.User, error) {
	var existing models.User
	if err := s.DB.Where("username = ?", username).First(&existing).Error; err == nil {
		return nil, errors.New("username already exists")
	} else if err != gorm.ErrRecordNotFound {
		return nil, err
	}

	user := models.User{
		ID:       uuid.New(), // generate a new UUID
		Username: username,
		Score:    0,
	}
	if err := s.DB.Create(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func (s *AuthService) GetProfile(userID uuid.UUID) (*models.User, error) {
	var user models.User
	if err := s.DB.First(&user, "id = ?", userID).Error; err != nil {
		return nil, err
	}
	return &user, nil
}
