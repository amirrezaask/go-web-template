package mb

import (
	"template/application"

	"github.com/nats-io/stan.go"
)

type natsStreaming struct {
	Conn stan.Conn
}

func (n *natsStreaming) Subscribe(topic string, handler func(message []byte)) error {
	_, err := n.Conn.Subscribe(topic, func(m *stan.Msg) {
		handler(m.Data)
	})
	if err != nil {
		return err
	}
	return nil
}

func (n *natsStreaming) Publish(topic string, message []byte) error {
	return n.Conn.Publish(topic, message)
}

type NatsStreamingConfig struct {
	Options   []stan.Option
	ClusterID string
	ClientID  string
}

func newNatsStreaming(c *NatsStreamingConfig) (application.MessageBroker, error) {
	conn, err := stan.Connect(c.ClusterID, c.ClientID, c.Options...)
	if err != nil {
		return nil, err
	}
	return &natsStreaming{conn}, nil
}
