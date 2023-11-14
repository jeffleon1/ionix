package entities

import (
	"time"

	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"
)

type Vaccination struct {
	gorm.Model
	Name   string    `json:"name" validate:"required,min=3,max=50"`
	DrugID int       `json:"drug_id" validate:"gte=0"`
	Dose   int       `json:"dose" validate:"gt=0"`
	Date   time.Time `json:"date" validate:"required"`
}

func (v *Vaccination) Validate() error {
	validate := validator.New()
	return validate.Struct(v)
}
