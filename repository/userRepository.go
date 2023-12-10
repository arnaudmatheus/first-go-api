package repository

import (
	"example/first-api/schemas"

	"gorm.io/gorm"
)

type UserRepository struct {
	Db *gorm.DB
}

func InitializeUserRepository(db *gorm.DB) UserRepositoryInterface {
	return &UserRepository{
		Db: db,
	}
}

func (t *UserRepository) CreateUser(person schemas.Person) error {
	if err := t.Db.Create(&person).Error; err != nil {
		return err
	}
	return nil
}
