package messaging

import (
	"encoding/json"
	"log"

	"github.com/nats-io/nats.go"

	domainevent "github.com/glaubersc/ecosystem/services/service-template/internal/domain/event"
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
	_, err := c.client.JS.Subscribe(c.subject, func(msg *nats.Msg) {
		var evt domainevent.Envelope

		if err := json.Unmarshal(msg.Data, &evt); err != nil {
			log.Printf("invalid event payload: %v", err)
			return
		}

		if err := c.handler(evt); err != nil {
			log.Printf("event handling failed: %v", err)
			return
		}

		// Acknowledge only on success (JetStream best practice)
		_ = msg.Ack()
	})

	return err
}
