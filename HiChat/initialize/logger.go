package initialize

import (
	"go.uber.org/zap"
	"log"
)

func InitLogger() {
	//初始化日志
	logger, err := zap.NewDevelopment()
	if err != nil {
		log.Fatal("日志初始化失败", err.Error())
	}
	//使用全局logger
	zap.ReplaceGlobals(logger)
}
