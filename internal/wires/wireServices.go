//go:build wireinject
// +build wireinject

// The build tag makes sure the stub is not built in the final build.

package wires

import (
	"aCupOfGin/internal/repos/userRepo"
	"aCupOfGin/internal/services/userService"
	"github.com/google/wire"
)

func InitUserService(ur userRepo.InterfaceUserRepo) *userService.ImplementUserService {
	wire.Build(userService.NewUserService)
	return nil
}
