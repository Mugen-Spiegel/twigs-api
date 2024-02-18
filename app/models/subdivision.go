package models

import (
	"time"

	"github.com/google/uuid"
)

type Subdivision struct {
	UUID       uuid.UUID `json:"uuid" gorm:"primary_key"`
	Name       string    `json:"name"`
	Barangay   string    `json:"barangay"`
	City       string    `json:"city"`
	PostalCode string    `json:"postal_code"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
}
