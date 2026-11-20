package fix_tester

import "github.com/funtimecoding/soil/pkg/lint/concern"

func (o *Tester) Planned(
	entries []*concern.Concern,
	path string,
	message string,
) {
	o.t.Helper()
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
			o.t.Errorf("concern %q is not a plain suggestion", message)
		}

		return
	}

	o.t.Errorf(
		"expected suggestion {path: %q, text: %q} not found, found: %q",
		path,
		message,
		found,
	)
}
