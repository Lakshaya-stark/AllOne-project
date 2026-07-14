package models

import (
	"time"

	"github.com/google/uuid"
)

type Device struct {
	ID         uuid.UUID
	UserID     uuid.UUID
	Name       string
	Platform   string
	DeviceType string
	PublicKey  string
	LastSeen   *time.Time
	CreatedAt  time.Time
}