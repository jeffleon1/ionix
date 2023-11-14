package db

import (
	"log"

	"github.com/jeffleon1/ionix/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

// connect database
func Connect(cfg *config.Config) *gorm.DB {
	db, err := gorm.Open(postgres.Open(cfg.ConnectionString()), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Silent),
	})
	if err != nil {
		log.Fatal(err)
	}
	return db
}
