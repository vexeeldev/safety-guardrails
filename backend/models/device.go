package models

import (
	"time"

	"github.com/google/uuid"
)

type Device struct {
	ID        uuid.UUID `gorm:"type:uuid;primaryKey;default:gen_random_uuid()"`
	UserID    uuid.UUID `gorm:"type:uuid;not null"`
	Name      string    `gorm:"not null"`
	Type      string    `gorm:"not null"` // e.g. "android", "desktop"
	Status    string    `gorm:"not null"` // e.g. "online", "offline"
	LastSeen  time.Time
	CreatedAt time.Time
	UpdatedAt time.Time
}
