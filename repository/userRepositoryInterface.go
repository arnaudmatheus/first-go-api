package repository

import "example/first-api/schemas"

type UserRepositoryInterface interface {
	CreateUser(person schemas.Person) error
}
