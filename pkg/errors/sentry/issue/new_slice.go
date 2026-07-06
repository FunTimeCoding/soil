package issue

import "github.com/funtimecoding/soil/pkg/errors/sentry/basic/response"

func NewSlice(v []response.Issue) []*Issue {
	var result []*Issue

	for _, e := range v {
		result = append(result, New(&e))
	}

	return result
}
