package main

import (
	_ "example/first-api/docs"

	"example/first-api/config"
	"example/first-api/handler"
	"example/first-api/router"
)

// @title Meu Primeiro CRUD em Go | HunCoding
// @version 1.0
// @description API for crud operations on users
// @host localhost:8080
// @BasePath /
// @schemes http
// @license MIT
func main() {
	database := config.Init()
	userRequest := handler.InitializeUserRequest(database)
	router.InitializeRouter(userRequest)
}
