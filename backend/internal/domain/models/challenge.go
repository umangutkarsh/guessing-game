package models

import (
	"time"

	"github.com/google/uuid"
)

type Challenge struct {
	ID            uuid.UUID   `gorm:"type:uuid;primaryKey"`
	InviterUserID uuid.UUID      `gorm:"type:uuid"`
	Token         string    `gorm:"unique;not null"`
	InviteLink    string    `gorm:"not null"`
	CreatedAt     time.Time
}
