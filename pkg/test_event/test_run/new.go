package test_run

import "github.com/funtimecoding/soil/pkg/test_event"

func New(events []*test_event.Event) *Run {
	return &Run{events: events}
}
