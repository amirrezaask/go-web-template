package mb

type MessageBroker interface {
	Subscribe(topic string, handler func(message []byte) error) error
	Publish(message []byte) error
}

func NewMessageBroker() (MessageBroker, error) {
	return nil, nil
}
