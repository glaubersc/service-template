package event

import "github.com/your-org/service-template/internal/domain/event"

type Handler interface {
	EventType() string
	Handle(evt event.Envelope) error
}
