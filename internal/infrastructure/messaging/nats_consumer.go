package messaging

import (
	"encoding/json"
	"log"

	domainevent "github.com/your-org/service-template/internal/domain/event"
)

type Consumer struct {
	client  *NATSClient
	handler func(domainevent.Envelope) error
	subject string
}

func NewConsumer(
	client *NATSClient,
	subject string,
	handler func(domainevent.Envelope) error,
) *Consumer {
	return &Consumer{
		client:  client,
		subject: subject,
		handler: handler,
	}
}

func (c *Consumer) Start() error {
	_, err := c.client.JS.Subscribe(c.subject, func(msg []byte) {
		var evt domainevent.Envelope

		if err := json.Unmarshal(msg, &evt); err != nil {
			log.Printf("invalid event payload: %v", err)
			return
		}

		if err := c.handler(evt); err != nil {
			log.Printf("event handling failed: %v", err)
			return
		}
	})

	return err
}
