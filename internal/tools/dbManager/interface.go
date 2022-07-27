package dbManager

import "gorm.io/gorm"

type InterfaceDBManger interface {
	Init(Dialector gorm.Dialector, DBMaxIdleConns int, DBMaxOpenConns int, ConnMaxLifeTimeMinutes int)
	IsConnected() bool
	ProvideConnection() *gorm.DB
}
