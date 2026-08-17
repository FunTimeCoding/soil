package forward

import "github.com/funtimecoding/soil/pkg/opnsense/response"

func NewSlice(v []response.Forward) []*Forward {
	var result []*Forward

	for _, e := range v {
		result = append(result, New(e))
	}

	return result
}
