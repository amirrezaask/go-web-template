package monitoring

import (
	"app/config"

	"github.com/sirupsen/logrus"
)

var logLevelMapper = map[string]logrus.Level{
	"trace": logrus.TraceLevel,
	"debug": logrus.DebugLevel,
	"info":  logrus.InfoLevel,
	"warn":  logrus.WarnLevel,
	"error": logrus.ErrorLevel,
	"fatal": logrus.FatalLevel,
}

var loggerInstance *logrus.Logger

func Logger() *logrus.Logger {
	if loggerInstance != nil {
		return loggerInstance
	}
	config.C.SetDefault("log.level", "debug")
	level := config.C.GetString("log.level")
	loggerInstance = logrus.New()
	loggerInstance.SetLevel(logLevelMapper[level])
	loggerInstance.SetFormatter(&logrus.JSONFormatter{})
	return loggerInstance
}
