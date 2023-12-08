package main

import (
	"example/first-api/config"
	"example/first-api/router"
)

func main() {
	config.Init()
	router.InitializeRouter()
}
