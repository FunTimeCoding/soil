package unclosed_resource

import "github.com/funtimecoding/soil/pkg/lint/fact"

func (s *Summaries) holdsAll(conditions []fact.Condition) bool {
	for i := range conditions {
		if !s.holds(&conditions[i]) {
			return false
		}
	}

	return true
}
