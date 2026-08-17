package rule

import "github.com/funtimecoding/soil/pkg/opnsense/response"

func NewSlice(v []response.Rule) []*Rule {
	var result []*Rule

	for _, e := range v {
		result = append(result, New(e))
	}

	return result
}
