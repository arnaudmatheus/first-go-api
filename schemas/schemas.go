package schemas

import (
	"time"

	"gorm.io/gorm"
)

type Person struct {
	gorm.Model
	Name     string
	Cpf      string
	Location string
	Salary   int64
}

type PersonResponse struct {
	ID        uint      `json:"id"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
	DeletedAt time.Time `json:"deletedAt,omitempty"`
	Name      string    `json:"name"`
	Cpf       string    `json:"cpf"`
	Location  string    `json:"location"`
	Salary    int64     `json:"salary"`
}

type PersonRequest struct {
	Name     string `json:"name"`
	Cpf      string `json:"cpf"`
	Location string `json:"location"`
	Salary   int64  `json:"salary"`
}
