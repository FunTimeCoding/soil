package service

import "github.com/funtimecoding/soil/pkg/lint/output"

func hasUnfixed(r *output.Results) bool {
	for _, c := range r.Entries {
		if !c.Fixed {
			return true
		}
	}

	return false
}
