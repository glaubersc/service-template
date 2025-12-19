package event

import "time"

type Envelope struct {
	EventID     string    `json:"eventId"`
	EventType   string    `json:"eventType"`
	AggregateID string    `json:"aggregateId"`
	OccurredAt  time.Time `json:"occurredAt"`
	Producer    string    `json:"producer"`
	TraceID     string    `json:"traceId"`
	Payload     any       `json:"payload"`
}
