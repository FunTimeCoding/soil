package state

import "github.com/funtimecoding/soil/pkg/opnsense/response"

func NewSlice(v []response.State) []*State {
	var result []*State

	for _, e := range v {
		result = append(result, New(e))
	}

	return result
}
