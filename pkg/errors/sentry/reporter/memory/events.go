package memory

import (
	"github.com/funtimecoding/soil/pkg/errors/types/reported_event"
	"slices"
)

func (m *Memory) Events() []*reported_event.Event {
	m.mutex.Lock()
	defer m.mutex.Unlock()

	return slices.Clone(m.events)
}
