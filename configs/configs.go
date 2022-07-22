package configs

import (
	"github.com/spf13/viper"
	"reflect"
)

var ConfigSet *Config

type Config struct {
	App *AppConf
	Db  *DbConf
}

type AppConf struct {
	Project   string `mapstructure:"APP_PROJECT"`
	UrlPrefix string `mapstructure:"APP_URL_PREFIX"`
	Port      string `mapstructure:"APP_PORT"`
	Version   string `mapstructure:"APP_VERSION"`
	GinMode   string `mapstructure:"APP_GIN_MODE"`
}

func NewAppConf() *AppConf {
	return &AppConf{
		Project:   "HelloGin",
		UrlPrefix: "/",
		Port:      "9000",
		Version:   "v0.0.1",
		GinMode:   "release",
	}
}

type DbConf struct {
	Host                   string `mapstructure:"DB_HOST"`
	UserName               string `mapstructure:"DB_USER_NAME"`
	Password               string `mapstructure:"DB_PWD"`
	DbName                 string `mapstructure:"DB_DB_NAME"`
	Port                   string `mapstructure:"DB_PORT"`
	MaxIdleConnection      int    `mapstructure:"DB_MAX_IDLE_CONN"`
	MaxOpenConnection      int    `mapstructure:"DB_MAX_OPEN_CONN"`
	ConnMaxLifetimeMinutes int    `mapstructure:"DB_CONN_MAX_LIFE_MIN"`
}

func NewDbConf() *DbConf {
	return &DbConf{
		Host:                   "localhost",
		UserName:               "test_user",
		Password:               "test-pws",
		DbName:                 "user_profile",
		Port:                   "3306",
		MaxIdleConnection:      1,
		MaxOpenConnection:      2,
		ConnMaxLifetimeMinutes: 60,
	}
}

func readCustomEnv() {
	viper.AddConfigPath("./configs")
	viper.SetConfigName(".env")
	viper.SetConfigType("env")

	errReadInConfig := viper.ReadInConfig()
	if errReadInConfig != nil {
		if reflect.TypeOf(errReadInConfig).Kind() != reflect.TypeOf(viper.ConfigFileNotFoundError{}).Kind() {
			panic("Error on Read Custom CONFIG" + errReadInConfig.Error())
		}
	}

}

func init() {
	readCustomEnv()
	viper.AutomaticEnv()

	appConfig := NewAppConf()
	err := viper.Unmarshal(appConfig)
	if err != nil {
		panic("unable to decode into config struct, %v")
	}

	dbConfig := NewDbConf()
	err = viper.Unmarshal(dbConfig)
	if err != nil {
		panic("unable to decode into config struct, %v")
	}

	ConfigSet = &Config{
		App: appConfig,
		Db:  dbConfig,
	}

}
