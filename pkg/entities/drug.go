package entities

import (
	"time"

	"gorm.io/gorm"
)

type Drug struct {
	gorm.Model
	Name        string    `json:"name" validate:"required,min=3,max=50"`
	Approved    bool      `json:"approved" validate:"required"`
	MinDose     int       `json:"min_dose" validate:"gt=0"`
	MaxDose     int       `json:"max_dose" validate:"gt=0"`
	AvailableAt time.Time `json:"available_at" validate:"required"`
}
