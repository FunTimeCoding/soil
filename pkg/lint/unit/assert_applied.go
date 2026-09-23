package unit

import (
	"github.com/funtimecoding/soil/pkg/lint/concern"
	"testing"
)

func assertApplied(
	t *testing.T,
	entries []*concern.Concern,
	path string,
	message string,
) {
	t.Helper()

	for _, c := range entries {
		if c.Path == path && c.Text == message && c.Fixed {
			return
		}
	}

	t.Errorf(
		"expected applied concern {path: %q, text: %q} not found",
		path,
		message,
	)
}
