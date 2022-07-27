package userService

import (
	"aCupOfGin/internal/repos/userRepo"
	"fmt"
	"github.com/fatih/structs"
)

type ImplementUserService struct {
	UserRepo userRepo.InterfaceUserRepo
}

func (userService *ImplementUserService) CreateUser(vendor string, account string,
	accountType string, hashedPwd string, name string) bool {

	status, err := userService.UserRepo.CreateUser(vendor, account, accountType, hashedPwd, name)
	if err != nil {
		fmt.Println(err.Error())
	}
	return status
}

func (userService *ImplementUserService) GetAllUsers() []map[string]interface{} {
	users := userService.UserRepo.GetAllUsers()
	output := make([]map[string]interface{}, 0, len(users))
	for _, user := range users {
		output = append(output, structs.Map(user))
	}
	return output
}

func (userService *ImplementUserService) GetUser(id int) map[string]interface{} {
	output := map[string]any{}
	user := userService.UserRepo.GetUser(id)
	if user == nil {
		return output
	}
	output = structs.Map(user)
	return output
}

func (userService *ImplementUserService) DeleteUser(id int) bool {
	status, err := userService.UserRepo.DeleteUser(id)
	if err != nil {
		fmt.Println(err.Error())
	}
	return status
}

func (userService *ImplementUserService) UpdateUser(id int, name string) bool {
	status, err := userService.UserRepo.UpdateUser(id, name)
	if err != nil {
		fmt.Println(err.Error())
	}
	return status
}

func NewUserService(ur userRepo.InterfaceUserRepo) *ImplementUserService {
	return &ImplementUserService{UserRepo: ur}
}
