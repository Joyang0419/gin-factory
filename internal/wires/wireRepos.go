//go:build wireinject
// +build wireinject

// The build tag makes sure the stub is not built in the final build.

package wires

import (
	"aCupOfGin/internal/repos/userRepo"
	"aCupOfGin/internal/tools/dbManager"
	"github.com/google/wire"
)

func InitGORMUserRepo(dm dbManager.InterfaceDBManger) *userRepo.GormUserRepo {
	wire.Build(userRepo.NewGormUserRepo)
	return nil
}

func InitCSVUserRepo(dm dbManager.InterfaceDBManger) *userRepo.CSVUserRepo {
	wire.Build(userRepo.NewCSVUserRepo)
	return nil
}
