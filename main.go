// 项目启动
package main

import (
	"go.uber.org/zap"
	"panda/initialize"
)

func main() {

	zapLogger, _ := zap.NewProduction()

	defer func(zapLogger *zap.Logger) {
		_ = zapLogger.Sync()
	}(zapLogger)

	// 创建自定义 GORM 记录器
	customLogger := &initialize.CustomLogger{ZapLogger: zapLogger}

	// 使用自定义记录器初始化 GORM
	initialize.CustomGorm(customLogger)

}
