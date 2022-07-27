//go:build wireinject
// +build wireinject

// The build tag makes sure the stub is not built in the final build.

package wires

import (
	"aCupOfGin/internal/controllers/userController"
	"aCupOfGin/internal/services/userService"
	"github.com/google/wire"
)

func InitUserController(us userService.InterfaceUserService) *userController.ImplementUserController {
	wire.Build(userController.NewUserController)
	return nil
}
