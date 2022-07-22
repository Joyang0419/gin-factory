package controller

import (
	"aCupOfGin/internal/common/rsp"
	"aCupOfGin/internal/entity"
	"aCupOfGin/internal/logger"
	"aCupOfGin/internal/service"
	"github.com/gin-gonic/gin"
)

func CreateUser(c *gin.Context) {

	var user entity.User
	c.BindJSON(&user)

	err := service.CreateUser(&user)

	if err != nil {
		rsp.Error(c, err.Error())
	} else {
		rsp.Success(c, "success", user)
	}
}

// GetUser swagger annotation
// @Summary Get a User
// @Schemes
// @Description Get a User by ID
// @Tags users
// @Accept json
// @Produce json
// @Param id path int true "User ID"
// @Success 200 {string} string	"ok"
// @Router /user/users/{id} [get]
func GetUser(c *gin.Context) {
	logger.Logger.Debug("demo how to use logger")
	logger.Logger.Info("demo how to use logger")

	id, ok := c.Params.Get("id")
	if !ok {
		rsp.Error(c, "invalid id")
	}
	todoList, err := service.GetUserById(id)
	if err != nil {
		rsp.Error(c, err.Error())
	} else {
		rsp.Success(c, "success", todoList)
	}
}

func GetUserList(c *gin.Context) {
	todoList, err := service.GetAllUser()
	if err != nil {
		rsp.Error(c, err.Error())
	} else {
		rsp.Success(c, "success", todoList)
	}
}

func UpdateUser(c *gin.Context) {
	id, ok := c.Params.Get("id")
	if !ok {
		rsp.Error(c, "invalid id")
	}
	user, err := service.GetUserById(id)
	if err != nil {
		rsp.Error(c, err.Error())
		return
	}

	oriPwd := user.Password
	c.BindJSON(&user)
	newPwd := user.Password

	var isPwdUpdated bool
	if oriPwd == newPwd {
		isPwdUpdated = false
	} else {
		isPwdUpdated = true
	}

	if err = service.UpdateUser(user, isPwdUpdated); err != nil {
		rsp.Error(c, err.Error())
	} else {
		rsp.Success(c, "success", user)
	}
}

func DeleteUserById(c *gin.Context) {
	id, ok := c.Params.Get("id")
	if !ok {
		rsp.Error(c, "invalid id")
	}
	if err := service.DeleteUserById(id); err != nil {
		rsp.Error(c, err.Error())
	} else {
		rsp.Success(c, "success", id)
	}
}
