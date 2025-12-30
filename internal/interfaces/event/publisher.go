package event

import "github.com/glaubersc/ecosystem/services/service-template/internal/domain/event"

type Publisher interface {
	Publish(evt event.Envelope) error
}
