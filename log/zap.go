package log

import "go.uber.org/zap"

func newZap(level LogLevel) ILogger {
	switch level {
	case DebugLevel:
		l, err := zap.NewDevelopment()
		if err != nil {
			panic(err)
		}
		sl := l.Sugar()
		return sl
	case ErrorLevel:
		l, err := zap.NewProduction()
		if err != nil {
			panic(err)
		}
		sl := l.Sugar()
		return sl
	default:
		l, err := zap.NewDevelopment()
		if err != nil {
			panic(err)
		}
		sl := l.Sugar()
		return sl
	}
}
