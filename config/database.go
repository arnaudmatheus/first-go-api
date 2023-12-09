package config

import (
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

	if err != nil {
		fmt.Println("Erro ao conectar em banco")
	}

	err = db.AutoMigrate(&schemas.Person{})
	if err != nil {
		fmt.Println("Erro ao criar tabela")
		return nil, err
	}

	return db, err
}
