package handler

import "github.com/gin-gonic/gin"

type UserHandlerInterface interface {
	CreateUser(*gin.Context)
}
