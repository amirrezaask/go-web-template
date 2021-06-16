package logger

import (
	"go.uber.org/zap"
)

var Dev = zap.NewDevelopmentConfig
var Prod = zap.NewProductionConfig

func NewZap(configBuilder func() zap.Config) (Logger, error) {
    logger, err := configBuilder().Build()
    if err != nil {
      return nil, err
    }
    sl := logger.Sugar()
    return sl, nil
}

