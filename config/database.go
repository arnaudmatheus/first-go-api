package config

import (
	"example/first-api/helper"
	"example/first-api/schemas"
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

const (
	PostgresDriver = "postgres"
	User           = "boms"
	Host           = "localhost"
	Port           = "5432"
	Password       = "boms"
	DbName         = "go-db"
)

var DataSourceName = fmt.Sprintf("host=%s port=%s user=%s "+
	"password=%s dbname=%s sslmode=disable", Host, Port, User, Password, DbName)

func databaseConnectionInitializer() (*gorm.DB, error) {
	fmt.Println("Iniciando conexão com banco")
	db, err := gorm.Open(postgres.Open(DataSourceName), &gorm.Config{})

	helper.Error(err)

	err = db.AutoMigrate(&schemas.Person{})

	helper.Error(err)

	return db, err
}
