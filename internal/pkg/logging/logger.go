package logging

type logger interface {
	Info(args ...interface{})
	Infof(format string, args ...interface{})
	Infow(msg string, keysAndValues ...interface{})

	Warn(args ...interface{})
	Warnf(format string, args ...interface{})
	Warnw(msg string, keysAndValues ...interface{})

	Error(args ...interface{})
	Errorf(format string, args ...interface{})
	Errorw(msg string, keysAndValues ...interface{})

	Fatal(args ...interface{})
	Fatalf(format string, args ...interface{})
	Fatalw(msg string, keysAndValues ...interface{})
}

type Logger struct {
	logger
}

//global logger
var gLogger *Logger

func Infof(format string, args ...interface{}) {
	gLogger.Infof(format, args...)
}

func Info(args ...interface{}) {
	gLogger.Info(args...)
}

func Infow(msg string, keysAndValues ...interface{}) {
	gLogger.Infow(msg, keysAndValues...)
}

func Warn(args ...interface{}) {
	gLogger.Warn(args...)
}

func Warnf(format string, args ...interface{}) {
	gLogger.Warnf(format, args...)
}
func Warnw(msg string, keysAndValues ...interface{}) {
	gLogger.Infow(msg, keysAndValues...)
}

func Errorw(msg string, keysAndValues ...interface{}) {
	gLogger.Infow(msg, keysAndValues...)
}

func Error(args ...interface{}) {
	gLogger.Error(args...)
}

func Errorf(format string, args ...interface{}) {
	gLogger.Errorf(format, args...)
}

func Fatal(args ...interface{}) {
	gLogger.Fatal(args...)
}

func Fatalf(format string, args ...interface{}) {
	gLogger.Fatalf(format, args...)
}

func Fatalw(msg string, keysAndValues ...interface{}) {
	gLogger.Infow(msg, keysAndValues...)
}
