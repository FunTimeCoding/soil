package alias

import "github.com/funtimecoding/soil/pkg/opnsense/response"

func NewSlice(v []response.Alias) []*Alias {
	var result []*Alias

	for _, e := range v {
		result = append(result, New(e))
	}

	return result
}
