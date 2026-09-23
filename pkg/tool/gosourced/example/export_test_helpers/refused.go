package export_test_helpers

import (
	"github.com/funtimecoding/soil/pkg/lint/output"
	"strings"
)

func refused(r *output.Results) string {
	if r == nil {
		return ""
	}

	var reasons []string

	for _, c := range r.Entries {
		if !c.Fixed && !c.Planned {
			reasons = append(reasons, c.Text)
		}
	}

	return strings.Join(reasons, "; ")
}
