package list

import "github.com/funtimecoding/soil/pkg/linkace/response"

func NewSlice(
	v []response.List,
	host string,
) []*List {
	var result []*List

	for _, e := range v {
		result = append(result, New(e, host))
	}

	return result
}
