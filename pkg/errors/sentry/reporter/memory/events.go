package memory

import "github.com/funtimecoding/soil/pkg/errors/types/reported_event"

func (m *Memory) Events() []*reported_event.Event {
	return m.events
}
