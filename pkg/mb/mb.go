package mb

import "fmt"

type MessageBroker interface {
	Subscribe(topic string, handler func(message []byte)) error
	Publish(topic string, message []byte) error
}

func NewMessageBroker(backend string, opts interface{}) (MessageBroker, error) {
	switch backend {
	case "nats":
		return newNats(opts.(*NatsMbConfig))
	}
	return nil, fmt.Errorf("no backend matched for %v", backend)
}
