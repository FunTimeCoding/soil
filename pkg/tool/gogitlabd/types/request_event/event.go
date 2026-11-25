package request_event

import "github.com/funtimecoding/soil/pkg/watch/digest/event"

type Event struct {
	Event  *event.Event
	Writer string
}
