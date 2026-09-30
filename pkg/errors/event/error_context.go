package event

import "github.com/funtimecoding/soil/pkg/errors/constant"

func (e *EventError) ErrorContext() (string, map[string]any) {
	return constant.Event, map[string]any{
		constant.Type: e.Type,
		constant.Raw:  e.Raw,
	}
}
