package service_tester

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/lint/output"
)

func ConcernText(r *output.Results) []string {
	var result []string

	for _, c := range r.Entries {
		result = append(result, fmt.Sprintf("%s: %s", c.Path, c.Text))
	}

	return result
}
