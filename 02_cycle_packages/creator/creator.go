package creator

import "github.com/itiky/practicum-examples/02_cycle_packages/eventservice"

type EventCreator struct {
}

func New() *EventCreator {
	return &EventCreator{}
}

func (*EventCreator) Create() eventservice.Event {
	return eventservice.Event{Id: 1}
}


