package unit_test

import (
	"github.com/funtimecoding/soil/pkg/lint/concern"
	"testing"
)

func filterApplied(entries []*concern.Concern) []*concern.Concern {
	var r []*concern.Concern

	for _, c := range entries {
		if c.Fixed {
			r = append(r, c)
		}
	}

	return r
}

func filterBlocked(entries []*concern.Concern) []*concern.Concern {
	var r []*concern.Concern

	for _, c := range entries {
		if !c.Fixed {
			r = append(r, c)
		}
	}

	return r
}

func assertResult(
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

	t.Errorf(
		"expected concern {path: %q, text: %q} not found",
		path,
		message,
	)
}
