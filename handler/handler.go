package handler

import (
	"example/first-api/config"

	"gorm.io/gorm"
)

var database *gorm.DB

func InitializeHandler() {
	database = config.GetDatabase()
}
