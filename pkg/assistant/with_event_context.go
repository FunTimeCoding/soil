package assistant

import (
	"github.com/funtimecoding/soil/pkg/assistant/message"
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/errors/event"
)

func withEventContext(
	m *message.Message,
	v any,
) any {
	if m.Event == nil {
		return v
	}

	return event.New(m.Event.Type, string(m.Event.Raw), errors.From(v))
}
