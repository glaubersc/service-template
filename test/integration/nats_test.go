package integration

import (
	"testing"
	"time"

	"github.com/glaubersc/ecosystem/services/service-template/internal/domain/event"
	"github.com/glaubersc/ecosystem/services/service-template/internal/infrastructure/messaging"
)

func TestNATSPublish(t *testing.T) {
	client, err := messaging.Connect("nats://localhost:4222")
	if err != nil {
		t.Fatal(err)
	}
	defer client.Close()

	publisher := messaging.NewPublisher(client)

	evt := event.Envelope{
		EventID:     "test-id",
		EventType:   "TestEvent.v1",
		AggregateID: "agg-1",
		OccurredAt:  time.Now(),
		Producer:    "service-template",
		TraceID:     "trace-1",
		Payload:     map[string]string{"hello": "world"},
	}

	if err := publisher.Publish(evt); err != nil {
		t.Fatalf("failed to publish: %v", err)
	}
}
