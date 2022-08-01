package userController

import (
	"aCupOfGin/internal/common/rsp"
	"aCupOfGin/internal/entities"
	"aCupOfGin/internal/services/userService"
	"github.com/gin-gonic/gin"
	"strconv"
)

type ImplementUserController struct {
	UserService userService.InterfaceUserService
}

func NewUserController(us userService.InterfaceUserService) *ImplementUserController {
	return &ImplementUserController{UserService: us}
}

func (uc *ImplementUserController) CreateUser(c *gin.Context) {
	var user entities.UserEntity
	err := c.BindJSON(&user)

	if err != nil {
		return
	}

	createStatus := uc.UserService.CreateUser(user.Vendor, user.Account, user.AccountType, user.Password, user.Name)

	rsp.Success(c, "success", createStatus)

}

func (uc *ImplementUserController) GetUser(c *gin.Context) {
	id, ok := c.Params.Get("id")

	if ok == false {
		rsp.Error(c, "invalid id")
	}

	idInt, err := strconv.Atoi(id)

	if err != nil {
		rsp.Error(c, "invalid id")
	}

	user := uc.UserService.GetUser(idInt)
	rsp.Success(c, "success", user)
}

func (uc *ImplementUserController) GetUsers(c *gin.Context) {
	users := uc.UserService.GetAllUsers()
	rsp.Success(c, "success", users)
}

func (uc *ImplementUserController) UpdateUser(c *gin.Context) {
	id, ok := c.Params.Get("id")

	if ok == false {
		rsp.Error(c, "invalid id")
	}

	idInt, errId := strconv.Atoi(id)

	if errId != nil {
		rsp.Error(c, "invalid id")
	}

	var user entities.UserEntity
	err := c.BindJSON(&user)

	if err != nil {
		return
	}

	updateStatus := uc.UserService.UpdateUser(idInt, user.Name)
	rsp.Success(c, "success", updateStatus)

}

func (uc *ImplementUserController) DeleteUser(c *gin.Context) {
	id, ok := c.Params.Get("id")

	if ok == false {
		rsp.Error(c, "invalid id")
	}

	idInt, err := strconv.Atoi(id)

	if err != nil {
		rsp.Error(c, "invalid id")
	}

	deleteStatus := uc.UserService.DeleteUser(idInt)

	rsp.Success(c, "success", deleteStatus)

}
