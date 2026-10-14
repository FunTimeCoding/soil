package unclosed_resource

import "github.com/funtimecoding/soil/pkg/lint/fact"

func arranged(conditions []fact.Condition) bool {
	for _, c := range conditions {
		if c.Callee == "" {
			return true
		}
	}

	return false
}
