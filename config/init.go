package config

import (
	"fmt"

	"gorm.io/gorm"
)

var database *gorm.DB

func Init() error {
	var err error

	database, err = databaseConnectionInitializer()

	if err != nil {
		return fmt.Errorf("error initialize sqlite: %v", err)
	}

	return nil
}

func GetDatabase() *gorm.DB {
	return database
}
