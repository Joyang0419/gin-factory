package dao

import (
	"aCupOfGin/configs"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"time"
)

var SqlSession *gorm.DB

func InitMySql() {

	url := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		configs.ConfigSet.Db.UserName,
		configs.ConfigSet.Db.Password,
		configs.ConfigSet.Db.Host,
		configs.ConfigSet.Db.Port,
		configs.ConfigSet.Db.DbName,
	)

	SqlSession, _ = gorm.Open(mysql.Open(url), &gorm.Config{})

	sqlDB, _ := SqlSession.DB()
	sqlDB.SetMaxIdleConns(configs.ConfigSet.Db.MaxIdleConnection)
	sqlDB.SetMaxOpenConns(configs.ConfigSet.Db.MaxOpenConnection)
	sqlDB.SetConnMaxLifetime(time.Duration(configs.ConfigSet.Db.ConnMaxLifetimeMinutes) * time.Minute)
}
