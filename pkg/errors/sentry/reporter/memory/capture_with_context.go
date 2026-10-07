package memory

import "github.com/funtimecoding/soil/pkg/errors/types/reported_event"

func (m *Memory) CaptureWithContext(
	e error,
	_ string,
	context map[string]any,
) string {
	v := reported_event.New(e)
	v.Context = context
	m.events = append(m.events, v)

	return ""
}
