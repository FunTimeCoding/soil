package host

import "github.com/funtimecoding/soil/pkg/opnsense/response"

func NewSlice(v []response.Host) []*Host {
	var result []*Host

	for _, e := range v {
		result = append(result, New(e))
	}

	return result
}
