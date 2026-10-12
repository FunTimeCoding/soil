package toolchain

import (
	"github.com/funtimecoding/soil/pkg/lint/concern"
	"testing"
)

func assertPlanned(
	t *testing.T,
	entries []*concern.Concern,
	path string,
	message string,
) {
	t.Helper()
	var found []string

	for _, c := range entries {
		if c.Path != path {
			continue
		}

		if c.Text != message {
			found = append(found, c.Text)

			continue
		}

		if !c.Planned || c.Fixed {
			t.Errorf("concern %q is not a plain suggestion", message)
		}

		return
	}

	t.Errorf(
		"expected suggestion {path: %q, text: %q} not found, found: %q",
		path,
		message,
		found,
	)
}
