package models

import "github.com/google/uuid"

type User struct {
	ID       uuid.UUID `gorm:"type:uuid;primaryKey"`
	Username string `gorm:"unique;not null"`
	Score    int    `gorm:"default:0"`
}
