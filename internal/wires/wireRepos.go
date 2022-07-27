//go:build wireinject
// +build wireinject

// The build tag makes sure the stub is not built in the final build.

package wires

import (
	"aCupOfGin/internal/repos/userRepo"
	"aCupOfGin/internal/tools/dbManager"
	"github.com/google/wire"
)

func InitUserRepo(dm dbManager.InterfaceDBManger) *userRepo.GormUserRepo {
	wire.Build(userRepo.NewGormUserRepo)
	return nil
}
