package log

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

func Logger(logger string, level LogLevel) ILogger {
	if loggerInstance != nil {
		return loggerInstance
	}
	switch logger {
	case LoggerBackendZap:
		return newZap(level)
	default:
		return newZap(level)
	}
}
