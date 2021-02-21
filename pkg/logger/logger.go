package logger

import "fmt"

//Logger defines logger
type Logger interface {
	Debugf(format string, args ...interface{})
	Warnf(format string, args ...interface{})
	Errorf(format string, args ...interface{})
	Fatalf(format string, args ...interface{})
}

//NewLogger creates a new logger with given backend.
func NewLogger(backend string) (Logger, error) {
	switch backend {
	case "zap-devel":
		return newZapDevel()
	case "zap-prod":
		return newZapProd()
	}
	return nil, fmt.Errorf("no logger backend matched for %s", backend)
}
