package memory

import "github.com/funtimecoding/soil/pkg/errors/types/reported_event"

type Memory struct {
	events []*reported_event.Event
}
