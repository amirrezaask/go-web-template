package mb

import (
	"fmt"
	"template/application"
)

func NewMessageBroker(backend string, opts interface{}) (application.MessageBroker, error) {
	switch backend {
	case "nats":
		return newNats(opts.(*NatsMbConfig))
	}
	return nil, fmt.Errorf("no backend matched for %v", backend)
}
