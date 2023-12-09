package main

import (
	"example/first-api/config"
	"example/first-api/handler"
	"example/first-api/router"
)

func main() {
	database := config.Init()
	userRequest := handler.InitializeUserRequest(database)
	router.InitializeRouter(userRequest)
}
