package mb

import (
	"github.com/nats-io/nats.go"
)

func NewNats(url string, options ...nats.Option) (*nats.Conn, error) {
	conn, err := nats.Connect(url, options...)
	if err != nil {
		return nil, err
	}
	return conn, nil
}
