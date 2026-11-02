package unclosed_resource

import "github.com/funtimecoding/soil/pkg/lint/fact"

func fieldRelease(
	field string,
	releases []fact.Release,
) (*fact.Condition, bool) {
	var result *fact.Condition
	found := false

	for i := range releases {
		r := &releases[i]

		if r.Field != field {
			continue
		}

		if r.Condition.Kind == "" {
			return nil, true
		}

		if !found {
			result = &r.Condition
			found = true
		}
	}

	return result, found
}
