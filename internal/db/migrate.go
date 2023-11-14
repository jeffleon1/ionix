package db

import (
	"github.com/jeffleon1/ionix/pkg/entities"
	"gorm.io/gorm"
)

func MigrateDB(db *gorm.DB) error {
	return db.AutoMigrate(&entities.Drug{}, &entities.Vaccination{}, &entities.User{})
}
