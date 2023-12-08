package router

import (
	"github.com/gin-gonic/gin"
)

func InitializeRouter() {
	router := gin.Default()

	setupRoutes(router)

	router.Run(":8080")
}
