//go:build wireinject
// +build wireinject

// The build tag makes sure the stub is not built in the final build.

package wires

import (
	"aCupOfGin/internal/tools/dbManager"
	"github.com/google/wire"
)

func InitGormDBManager(Dialector dbManager.DBDialector, DBMaxIdleConns dbManager.DBMaxIdleConns, DBMaxOpenConns dbManager.DBMaxOpenConns, ConnMaxLifeTimeMinutes dbManager.ConnMaxLifeTimeMinutes) *dbManager.GormDBManager {
	wire.Build(dbManager.NewGormDBManager, dbManager.NewGORMDBMSetting)
	return nil
}

func InitCSVDBManager(filename dbManager.FileName) *dbManager.CSVDBManager {
	wire.Build(dbManager.NewCSVDBManager, dbManager.NewCSVDBMSetting)
	return nil
}
