package router

import (
	_ "example/first-api/docs"
	"example/first-api/handler"

	"github.com/gin-gonic/gin"
)

// @title Meu Primeiro CRUD em Go | HunCoding
// @version 1.0
// @description API for crud operations on users
// @host localhost:8080
// @BasePath /
// @schemes http
// @license MIT
func InitializeRouter(handler handler.UserHandlerInterface) {
	router := gin.Default()

	setupRoutes(router, handler)

	router.Run(":8080")
}
