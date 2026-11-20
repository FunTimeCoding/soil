package fix_tester

import "github.com/funtimecoding/soil/pkg/lint/concern"

func (o *Tester) Result(
	entries []*concern.Concern,
	path string,
	message string,
) {
	o.t.Helper()

	for _, c := range entries {
		if c.Path == path && c.Text == message {
			return
		}
	}

	o.t.Errorf("expected concern {path: %q, text: %q} not found", path, message)
}
