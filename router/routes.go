package router

import (
	"example/first-api/handler"

	"github.com/gin-gonic/gin"
)

func setupRoutes(router *gin.Engine) {
	handler.InitializeHandler()
	basePath := "/api/v1"

	v1 := router.Group(basePath)
	{
		v1.POST("/person", handler.CreateUser)
	}

}
