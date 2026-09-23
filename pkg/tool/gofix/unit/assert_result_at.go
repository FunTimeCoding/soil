package unit

import (
	"github.com/funtimecoding/soil/pkg/lint/concern"
	"testing"
)

func assertResultAt(
	t *testing.T,
	entries []*concern.Concern,
	path string,
	line int,
	message string,
) {
	t.Helper()

	for _, c := range entries {
		if c.Path == path && c.Line == line && c.Text == message {
			return
		}
	}

	t.Errorf(
		"expected concern {path: %q, line: %d, text: %q} not found",
		path,
		line,
		message,
	)
}
