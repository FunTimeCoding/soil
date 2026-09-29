package note

import "github.com/funtimecoding/soil/pkg/linkace/response"

func NewSlice(v []response.Note) []*Note {
	var result []*Note

	for _, e := range v {
		result = append(result, New(e))
	}

	return result
}
