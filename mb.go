package main

//MessageBroker defines message broker
type MessageBroker interface {
	Subscribe(topic string, handler func(message []byte)) error
	Publish(topic string, message []byte) error
}
