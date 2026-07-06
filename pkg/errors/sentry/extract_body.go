package sentry

import (
	"errors"
	"github.com/funtimecoding/soil/pkg/face"
)

func extractBody(e error) []byte {
	if e == nil {
		return nil
	}

	var b face.BodyProvider

	if errors.As(e, &b) {
		return b.Body()
	}

	return nil
}
