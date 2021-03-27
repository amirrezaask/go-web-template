package mb

import (
	"github.com/nats-io/stan.go"
)


func NewNatsStreaming(clusterID string, clientID string, options []stan.Option) (stan.Conn, error) {
	conn, err := stan.Connect(clusterID, clientID, options...)
	if err != nil {
		return nil, err
	}
	return conn, nil
}
