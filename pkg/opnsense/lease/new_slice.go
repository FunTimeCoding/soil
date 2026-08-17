package lease

import "github.com/funtimecoding/soil/pkg/opnsense/response"

func NewSlice(v []response.Lease) []*Lease {
	var result []*Lease

	for _, e := range v {
		result = append(result, New(e))
	}

	return result
}
