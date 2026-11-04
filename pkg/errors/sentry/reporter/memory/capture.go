package memory

import (
	"context"
	"github.com/funtimecoding/soil/pkg/errors/types/reported_event"
	"github.com/getsentry/sentry-go"
)

func (m *Memory) Capture(
	x context.Context,
	e error,
) string {
	v := reported_event.New(e)

	if h := sentry.GetHubFromContext(x); h != nil {
		v.Trace = h.GetTraceparent()
	}

	m.mutex.Lock()
	defer m.mutex.Unlock()
	m.events = append(m.events, v)

	return m.identifier()
}
