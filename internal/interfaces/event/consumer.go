package event

import "github.com/glaubersc/ecosystem/services/service-template/internal/domain/event"

type Handler interface {
	EventType() string
	Handle(evt event.Envelope) error
}
