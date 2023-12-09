package handler

import (
	"example/first-api/schemas"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type UserRequest struct {
	Db *gorm.DB
}

func InitializeUserRequest(db *gorm.DB) UserHandlerInterface {
	return &UserRequest{
		Db: db,
	}
}

func (t *UserRequest) CreateUser(c *gin.Context) {
	var newUser schemas.PersonRequest

	if err := c.BindJSON(&newUser); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	newUserEntity := schemas.Person{
		Name:     newUser.Name,
		Cpf:      newUser.Cpf,
		Location: newUser.Location,
		Salary:   newUser.Salary,
	}

	if err := t.Db.Create(&newUserEntity).Error; err != nil {
		fmt.Println("Erro ao criar usuário")

		c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": "database erro"})
		return
	}

	c.IndentedJSON(http.StatusCreated, newUserEntity)

}
