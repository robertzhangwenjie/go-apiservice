package logging

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var Logger *zap.SugaredLogger

func InitLogger() {
	// 自定义logger cfg
	cfg := zap.NewProductionConfig()
	// 日志级别
	//cfg.Level = zap.NewAtomicLevelAt(zapcore.InfoLevel)
	// 自定义消息的编码配置
	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder

	cfg.Encoding = "json"
	cfg.OutputPaths = []string{
		"stderr",
		"user.log",
	}
	cfg.EncoderConfig = encoderConfig
	logger, _ := cfg.Build()

	// 使用自定义的logger替换内置logger
	zap.ReplaceGlobals(logger)
	Logger = zap.S()
}
