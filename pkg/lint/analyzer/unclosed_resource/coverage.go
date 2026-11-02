package unclosed_resource

import "github.com/funtimecoding/soil/pkg/lint/fact"

func coverage(
	required map[string]bool,
	releases []fact.Release,
) ([]fact.Condition, bool) {
	var result []fact.Condition

	for field := range required {
		condition, found := fieldRelease(field, releases)

		if !found {
			return nil, false
		}

		if condition != nil {
			result = append(result, *condition)
		}
	}

	return result, true
}
