package log

import (
	"app/config"
	"github.com/labstack/echo/v4"
	"go.uber.org/zap"
)

func LoggerMiddleware(h echo.HandlerFunc) echo.HandlerFunc {
	return func(ctx echo.Context) error {
		Logger().Debugf("%s\n", ctx.Request().RequestURI)
		err := h(ctx)
		if ctx.Response().Status > 499 {
			Logger().Errorf("%s\n", err)
		}
		Logger().Debugf("%s\n", err)
		return err
	}
}

var loggerInstance *zap.SugaredLogger

func Logger() *zap.SugaredLogger {
	if loggerInstance != nil {
		return loggerInstance
	}
	if config.AppEnv() == "dev" || config.AppEnv() == "debug" {
		l, err := zap.NewDevelopment()
		if err != nil {
			panic(err)
		}
		loggerInstance = l.Sugar()
		return loggerInstance
	} else {
		l, err := zap.NewProduction()
		if err != nil {
			panic(err)
		}
		loggerInstance = l.Sugar()
		return loggerInstance
	}
}
