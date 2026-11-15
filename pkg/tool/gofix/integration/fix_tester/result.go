package fix_tester

import (
	"github.com/funtimecoding/soil/pkg/lint/concern"
	"testing"
)

func Result(
	t *testing.T,
	entries []*concern.Concern,
	path string,
	message string,
) {
	t.Helper()

	for _, c := range entries {
		if c.Path == path && c.Text == message {
			return
		}
	}

	t.Errorf("expected concern {path: %q, text: %q} not found", path, message)
}
