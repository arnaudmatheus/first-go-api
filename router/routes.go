package router

import (
	_ "example/first-api/docs"

	"example/first-api/handler"

	"github.com/gin-gonic/gin"

	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title Meu Primeiro CRUD em Go | HunCoding
// @version 1.0
// @description API for crud operations on users
// @host localhost:8080
// @BasePath /
// @schemes http
// @license MIT
func setupRoutes(router *gin.Engine, handler handler.UserHandlerInterface) {
	basePath := "/api/v1"

	v1 := router.Group(basePath)
	{
		v1.POST("/person", handler.CreateUser)

	}

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

}
