package dbManager

import (
	"fmt"
	"gorm.io/gorm"
	"time"
)

type GormDBManager struct {
	SqlSession *gorm.DB
}

func (Manager *GormDBManager) Init(Dialector gorm.Dialector, DBMaxIdleConns int, DBMaxOpenConns int, ConnMaxLifeTimeMinutes int) {
	sqlSession, err := gorm.Open(Dialector, &gorm.Config{})
	Manager.SqlSession = sqlSession

	if err != nil {
		panic("DataBase Connection Failed")
	}

	sqlDB, _ := sqlSession.DB()
	sqlDB.SetMaxIdleConns(DBMaxIdleConns)
	sqlDB.SetMaxOpenConns(DBMaxOpenConns)
	sqlDB.SetConnMaxLifetime(time.Duration(ConnMaxLifeTimeMinutes) * time.Minute)
}

func (Manager *GormDBManager) IsConnected() bool {
	_, err := Manager.SqlSession.DB()
	if err != nil {
		fmt.Println("get db failed:", err)
		return false
	}
	return true
}

func (Manager *GormDBManager) ProvideConnection() *gorm.DB {
	return Manager.SqlSession
}

type DBDialector gorm.Dialector
type DBMaxIdleConns int
type DBMaxOpenConns int
type ConnMaxLifeTimeMinutes int

type DBMSetting struct {
	Dialector              DBDialector
	DBMaxIdleConns         DBMaxIdleConns
	DBMaxOpenConns         DBMaxOpenConns
	ConnMaxLifeTimeMinutes ConnMaxLifeTimeMinutes
}

func NewDBMSetting(dialector DBDialector, dbMaxIdleConns DBMaxIdleConns,
	dbMaxOpenConns DBMaxOpenConns, ConnMaxLifeTimeMinutes ConnMaxLifeTimeMinutes) *DBMSetting {
	return &DBMSetting{Dialector: dialector,
		DBMaxIdleConns:         dbMaxIdleConns,
		DBMaxOpenConns:         dbMaxOpenConns,
		ConnMaxLifeTimeMinutes: ConnMaxLifeTimeMinutes}
}

func NewGormDBManager(dbmSetting *DBMSetting) *GormDBManager {
	Manager := GormDBManager{}
	Manager.Init(
		dbmSetting.Dialector,
		int(dbmSetting.DBMaxIdleConns),
		int(dbmSetting.DBMaxOpenConns),
		int(dbmSetting.ConnMaxLifeTimeMinutes))

	return &Manager
}
