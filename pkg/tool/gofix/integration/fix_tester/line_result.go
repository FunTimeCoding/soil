package fix_tester

import "github.com/funtimecoding/soil/pkg/lint/concern"

func (o *Tester) LineResult(
	entries []*concern.Concern,
	path string,
	line int,
	message string,
) {
	o.t.Helper()

	for _, c := range entries {
		if c.Path == path && c.Line == line && c.Text == message {
			return
		}
	}

	o.t.Errorf(
		"expected concern {path: %q, line: %d, text: %q} not found",
		path,
		line,
		message,
	)
}
