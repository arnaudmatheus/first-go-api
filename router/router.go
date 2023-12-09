package router

import (
	"example/first-api/handler"

	"github.com/gin-gonic/gin"
)

func InitializeRouter(handler handler.UserHandlerInterface) {
	router := gin.Default()

	setupRoutes(router, handler)

	router.Run(":8080")
}
