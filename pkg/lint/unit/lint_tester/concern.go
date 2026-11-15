package lint_tester

import (
	"github.com/funtimecoding/soil/pkg/lint/concern"
	"testing"
)

func Concern(
	t *testing.T,
	fixed bool,
	path string,
	text string,
	entries []*concern.Concern,
) {
	t.Helper()

	for _, c := range entries {
		if c.Path == path && c.Text == text && c.Fixed == fixed {
			return
		}
	}

	t.Errorf(
		"expected concern {path: %q, text: %q, fixed: %t} not found",
		path,
		text,
		fixed,
	)
}
