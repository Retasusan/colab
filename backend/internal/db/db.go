// Package db provides database connectivity for the application.
package db

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Open() (*gorm.DB, error) {
	dsn := "host=localhost user=app password=app dbname=colab port=5432 sslmode=disable"

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	return db, nil
}
