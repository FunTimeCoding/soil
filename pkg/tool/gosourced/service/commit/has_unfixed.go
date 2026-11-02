package commit

import "github.com/funtimecoding/soil/pkg/lint/output"

func HasUnfixed(r *output.Results) bool {
	for _, c := range r.Entries {
		if !c.Fixed {
			return true
		}
	}

	return false
}
