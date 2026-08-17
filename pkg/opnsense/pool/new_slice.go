package pool

import "github.com/funtimecoding/soil/pkg/opnsense/response"

func NewSlice(v []response.Pool) []*Pool {
	var result []*Pool

	for _, e := range v {
		result = append(result, New(e))
	}

	return result
}
