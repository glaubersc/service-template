package event

import "github.com/your-org/service-template/internal/domain/event"

type Publisher interface {
	Publish(evt event.Envelope) error
}
