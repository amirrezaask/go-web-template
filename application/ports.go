package application

//Logger defines logger
type Logger interface {
	Debugf(format string, args ...interface{})
	Warnf(format string, args ...interface{})
	Errorf(format string, args ...interface{})
	Fatalf(format string, args ...interface{})
}

//MessageBroker defines message broker
type MessageBroker interface {
	Subscribe(topic string, handler func(message []byte)) error
	Publish(topic string, message []byte) error
}
