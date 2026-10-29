package memory

import "github.com/funtimecoding/soil/pkg/errors/types/reported_event"

func (m *Memory) CaptureException(e error) string {
	m.mutex.Lock()
	defer m.mutex.Unlock()
	m.events = append(m.events, reported_event.New(e))

	return ""
}
