package mb

import (
	"template/application"

	"github.com/nats-io/nats.go"
)

type natsMb struct {
	Conn *nats.Conn
}

func (n *natsMb) Subscribe(topic string, handler func(message []byte)) error {
	_, err := n.Conn.Subscribe(topic, func(m *nats.Msg) {
		handler(m.Data)
	})
	if err != nil {
		return err
	}
	return nil
}
func (n *natsMb) Publish(topic string, message []byte) error {
	return n.Conn.Publish(topic, message)
}

type NatsMbConfig struct {
	Url     string
	Options []nats.Option
}

func newNats(c *NatsMbConfig) (application.MessageBroker, error) {
	conn, err := nats.Connect(c.Url, c.Options...)
	if err != nil {
		return nil, err
	}
	return &natsMb{conn}, nil
}
