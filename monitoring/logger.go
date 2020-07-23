package monitoring

import (
	"app/config"
	"github.com/labstack/echo/v4"
	"go.uber.org/zap"
	"log"

	"github.com/sirupsen/logrus"
)
func LoggerMiddleware(h echo.HandlerFunc) echo.HandlerFunc {
	return func(ctx echo.Context) error {
		Logger().Debugln(ctx.Request().RequestURI)
		err := h(ctx)
		if ctx.Response().Status > 499 {
			Logger().Errorln(err)
		}
		Logger().Debugln(err)
		return err
	}
}
var logLevelMapper = map[string]logrus.Level{
	"trace": logrus.TraceLevel,
	"debug": logrus.DebugLevel,
	"info":  logrus.InfoLevel,
	"warn":  logrus.WarnLevel,
	"error": logrus.ErrorLevel,
	"fatal": logrus.FatalLevel,
}

var loggerInstance *logrus.Logger

func Logger() *zap.SugaredLogger {
	if loggerInstance != nil {
		return loggerInstance
	}
	level, err := config.C.GetString("log.level")
	if err != nil {
	    log.Fatalln(err)
	}
	l, err := zap.NewProduction()
	return l.Sugar()
}
