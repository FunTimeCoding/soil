package request_event

import "github.com/funtimecoding/soil/pkg/watch/digest/event"

func New(
	e *event.Event,
	writer string,
) *Event {
	return &Event{Event: e, Writer: writer}
}
