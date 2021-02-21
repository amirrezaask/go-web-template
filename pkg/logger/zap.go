package logger

import (
	"go.uber.org/zap"
)

func newZapProd() (Logger, error) {
	l, err := zap.NewProduction()
	if err != nil {
		return nil, err
	}
	sl := l.Sugar()
	return sl, nil
}

func newZapDevel() (Logger, error) {
	l, err := zap.NewDevelopment()
	if err != nil {
		return nil, err
	}
	sl := l.Sugar()
	return sl, nil
}
