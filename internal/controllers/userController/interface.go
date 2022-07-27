package userController

import "github.com/gin-gonic/gin"

type InterfaceUserController interface {
	CreateUser(c *gin.Context)
	GetUser(c *gin.Context)
	GetUsers(c *gin.Context)
	UpdateUser(c *gin.Context)
	DeleteUser(c *gin.Context)
}
