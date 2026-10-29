package memory

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/errors/types/reported_event"
)

func (m *Memory) Recover(v any) string {
	e, okay := v.(error)

	if !okay {
		e = fmt.Errorf("%v", v)
	}

	m.mutex.Lock()
	defer m.mutex.Unlock()
	m.events = append(m.events, reported_event.New(e))

	return ""
}
