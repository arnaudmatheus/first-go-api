package config

import (
	"example/first-api/helper"

	"gorm.io/gorm"
)

func Init() *gorm.DB {

	database, err := databaseConnectionInitializer()

	helper.Error(err)

	return database
}
