package database

import (
	"WebParkir/apps/api/internal/domain"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func ConnectPostgres(databaseURL string) (*gorm.DB, error) {

	db, err := gorm.Open(postgres.Open(databaseURL), &gorm.Config{})

	if err != nil {
		return nil, err
	}

	return db, nil
}

func AutoMigrate(db *gorm.DB) error {
	return db.AutoMigrate(
		&domain.User{},
		&domain.Vehicle{},
	)
}
