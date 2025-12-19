package messaging

import (
	"encoding/json"

	domainevent "github.com/your-org/service-template/internal/domain/event"
)

type NATSPublisher struct {
	client *NATSClient
}

func NewPublisher(client *NATSClient) *NATSPublisher {
	return &NATSPublisher{client: client}
}

func (p *NATSPublisher) Publish(evt domainevent.Envelope) error {
	data, err := json.Marshal(evt)
	if err != nil {
		return err
	}

	subject := evt.EventType

	_, err = p.client.JS.Publish(subject, data)
	return err
}
