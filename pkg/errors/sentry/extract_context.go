package sentry

import (
	"errors"
	"github.com/funtimecoding/soil/pkg/face"
)

func extractContext(e error) (string, map[string]any) {
	if e == nil {
		return "", nil
	}

	var p face.ContextProvider

	if errors.As(e, &p) {
		return p.ErrorContext()
	}

	return "", nil
}
