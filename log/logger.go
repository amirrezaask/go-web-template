package log

import "app/config"

type LogLevel int

const (
	DebugLevel LogLevel = iota + 1
	ErrorLevel
	FatalLevel
)

type LoggerBackend string

const (
	LoggerBackendZap    = "zap"
	LoggerBackendLogrus = "logrus"
)

type ILogger interface {
	Errorf(format string, args ...interface{})
	Fatalf(format string, args ...interface{})
	Debugf(format string, args ...interface{})
}

var loggerInstance ILogger

func Logger() ILogger {
	if loggerInstance != nil {
		return loggerInstance
	}
	switch config.Config.Logger.Type {
	case LoggerBackendZap:
		loggerInstance = newZap(LogLevel(config.Config.Logger.Level))
	default:
		loggerInstance = newZap(LogLevel(config.Config.Logger.Level))
	}
	return loggerInstance
}
