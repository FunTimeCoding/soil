package tag

import "github.com/funtimecoding/soil/pkg/linkace/response"

func NewSlice(
	v []response.Tag,
	host string,
) []*Tag {
	var result []*Tag

	for _, e := range v {
		result = append(result, New(e, host))
	}

	return result
}
