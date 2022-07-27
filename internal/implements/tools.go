package implements

import (
	"aCupOfGin/configs"
	"aCupOfGin/internal/tools/dbManager"
	"aCupOfGin/internal/wires"
	"fmt"
	"gorm.io/driver/mysql"
)

var sqlUrl = fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local&timeout=5s",
	configs.ConfigSet.Db.UserName,
	configs.ConfigSet.Db.Password,
	configs.ConfigSet.Db.Host,
	configs.ConfigSet.Db.Port,
	configs.ConfigSet.Db.DbName,
)

var (
	DBManager = wires.InitGormDBManager(
		mysql.Open(sqlUrl),
		dbManager.DBMaxIdleConns(configs.ConfigSet.Db.MaxIdleConnection),
		dbManager.DBMaxOpenConns(configs.ConfigSet.Db.MaxOpenConnection),
		dbManager.ConnMaxLifeTimeMinutes(configs.ConfigSet.Db.ConnMaxLifetimeMinutes))
)
