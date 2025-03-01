package models

import (
	"github.com/google/uuid"
	"gorm.io/datatypes"
)

type Destination struct {
	ID       uuid.UUID   `gorm:"type:uuid;primaryKey"`
	City     string         `gorm:"not null"`
	Country  string         `gorm:"not null"`
	Clues    datatypes.JSON `gorm:"type:jsonb"`
	FunFact  datatypes.JSON `gorm:"type:jsonb"`
	Trivia   datatypes.JSON `gorm:"type:jsonb"`
}
