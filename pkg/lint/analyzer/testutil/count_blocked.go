package testutil

import "github.com/funtimecoding/soil/pkg/lint/output"

func countBlocked(results *output.Results) int {
	var count int

	for _, c := range results.Entries {
		if !c.Fixed {
			count++
		}
	}

	return count
}
