package logger

import (
	"go.uber.org/zap"
	"template/application"
)

func NewZap() (application.Logger, error) {
	l, err := zap.NewProduction()
	if err != nil {
		return nil, err
	}
	sl := l.Sugar()
	return sl, nil
}

func NewZapDevel() (application.Logger, error) {
	l, err := zap.NewDevelopment()
	if err != nil {
		return nil, err
	}
	sl := l.Sugar()
	return sl, nil
}
