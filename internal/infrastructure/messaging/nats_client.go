package messaging

import (
	"time"

	"github.com/nats-io/nats.go"
)

type NATSClient struct {
	Conn *nats.Conn
	JS   nats.JetStreamContext
}

func Connect(url string) (*NATSClient, error) {
	nc, err := nats.Connect(
		url,
		nats.Timeout(5*time.Second),
		nats.RetryOnFailedConnect(true),
		nats.MaxReconnects(10),
	)
	if err != nil {
		return nil, err
	}

	js, err := nc.JetStream()
	if err != nil {
		return nil, err
	}

	return &NATSClient{
		Conn: nc,
		JS:   js,
	}, nil
}

func (c *NATSClient) Close() {
	c.Conn.Close()
}
