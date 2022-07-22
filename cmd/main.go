package main

import (
	"aCupOfGin/api/swag/docs"
	"aCupOfGin/configs"
	"aCupOfGin/internal/controller"
	"aCupOfGin/internal/dao"
	"aCupOfGin/internal/logger"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func setGinSwagger(r *gin.Engine) {
	docs.SwaggerInfo.BasePath = configs.ConfigSet.App.UrlPrefix
	docs.SwaggerInfo.Title = configs.ConfigSet.App.Project
	docs.SwaggerInfo.Version = configs.ConfigSet.App.Version

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
}

func setRouter(r *gin.Engine) {
	userGroup := r.Group("user")
	{
		userGroup.POST("/users", controller.CreateUser)
		userGroup.GET("/users", controller.GetUserList)
		userGroup.GET("/users/:id", controller.GetUser)
		userGroup.PUT("/users/:id", controller.UpdateUser)
		userGroup.DELETE("/users/:id", controller.DeleteUserById)
	}

}

func setLogger(r *gin.Engine) {
	cfg := zap.Config{
		Encoding:          "console",
		OutputPaths:       []string{"stderr"},
		DisableStacktrace: false,
		EncoderConfig: zapcore.EncoderConfig{
			MessageKey:       "message",
			TimeKey:          "time",
			StacktraceKey:    "stacktrace",
			EncodeTime:       zapcore.ISO8601TimeEncoder,
			LevelKey:         "level",
			EncodeLevel:      zapcore.CapitalLevelEncoder,
			ConsoleSeparator: "\t| ",
		},
	}

	var level zapcore.Level
	if configs.ConfigSet.App.GinMode == "debug" {
		level = zapcore.DebugLevel
	} else {
		level = zapcore.InfoLevel
	}
	cfg.Level = zap.NewAtomicLevelAt(level)

	pLogger, _ := cfg.Build()

	r.Use(cors.New(CorsConfig()))
	r.Use(logger.GinLogger(pLogger), logger.GinRecovery(pLogger, true))

	logger.Logger = pLogger
}

func CorsConfig() cors.Config {
	corsConf := cors.DefaultConfig()
	corsConf.AllowAllOrigins = true
	corsConf.AllowMethods = []string{"GET", "POST", "DELETE", "OPTIONS", "PUT"}
	corsConf.AllowHeaders = []string{"Authorization", "Content-Type", "Upgrade", "Origin",
		"Connection", "Accept-Encoding", "Accept-Language", "Host",
		"Access-Control-Request-Method", "Access-Control-Request-Headers"}
	return corsConf
}

func main() {
	dao.InitMySql()

	gin.SetMode(configs.ConfigSet.App.GinMode)

	r := gin.New()

	setLogger(r)
	setRouter(r)
	setGinSwagger(r)

	err := r.Run("0.0.0.0:" + configs.ConfigSet.App.Port)
	if err != nil {
		return
	}
}
