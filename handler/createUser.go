package handler

import (
	"example/first-api/repository"
	"example/first-api/schemas"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type UserRequest struct {
	userRepository repository.UserRepositoryInterface
}

func InitializeUserRequest(userRepository repository.UserRepositoryInterface) UserHandlerInterface {
	return &UserRequest{
		userRepository: userRepository,
	}
}

// CreateUser Creates a new user
// @Summary Create a new user
// @Description Create a new user with the provided user information
// @Tags Users
// @Accept json
// @Produce json
// @Param userRequest body schemas.PersonRequest true "User information for registration"
// @Success 200 {object} schemas.PersonResponse
// @Router /api/v1/person [post]
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

	if err := t.userRepository.CreateUser(newUserEntity); err != nil {
		fmt.Println(err)
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.IndentedJSON(http.StatusCreated, newUserEntity)

}
