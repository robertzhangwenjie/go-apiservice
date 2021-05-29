package logging

import (
	"github.com/natefinch/lumberjack"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"os"
)

// InitWithZapLogger 初始化全局logger
func InitWithZapLogger(filename string) {

	encoderConfig := getEncoderConfig()
	zapLogger := newZapLogger(encoderConfig, filename)

	// 使用自定义的logger替换内置logger
	zap.ReplaceGlobals(zapLogger)
	// 将SugaredLogger赋值给自定义的Logger内嵌的logger接口，达到解耦
	gLogger = &Logger{zap.S()}

}

func getEncoderConfig() zapcore.EncoderConfig {
	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder

	return encoderConfig
}

func newZapLogger(encodeConfig zapcore.EncoderConfig, filename string) *zap.Logger {

	// 初始化一个WriterSync，当写日志时会同步写入到对应的Writer中
	errorWriter := zapcore.AddSync(os.Stderr)
	normalWriter := zapcore.AddSync(os.Stdout)
	allWriter := zapcore.AddSync(&lumberjack.Logger{
		Filename:   filename,
		MaxSize:    100,
		MaxAge:     7,
		MaxBackups: 35,
		LocalTime:  false,
		Compress:   false,
	})

	// 定义Encoder用于格式化日志
	jsonEncoder := zapcore.NewJSONEncoder(encodeConfig)

	// 只处理error以上的log
	errorLevel := zap.LevelEnablerFunc(func(lv1 zapcore.Level) bool {
		return lv1 >= zapcore.ErrorLevel
	})
	// 处理error以下级别的log
	normalLevel := zap.LevelEnablerFunc(func(lvl zapcore.Level) bool {
		return lvl < zapcore.ErrorLevel
	})
	// 处理所有的info
	allLevel := zap.LevelEnablerFunc(func(lvl zapcore.Level) bool {
		return lvl >= zapcore.DebugLevel
	})

	// 创建Core
	errorCore := zapcore.NewCore(jsonEncoder, errorWriter, errorLevel)
	normalCore := zapcore.NewCore(jsonEncoder, normalWriter, normalLevel)
	allCore := zapcore.NewCore(jsonEncoder, allWriter, allLevel)
	core := zapcore.NewTee(errorCore, normalCore, allCore)

	return zap.New(core)

}
