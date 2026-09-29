package link

import "github.com/funtimecoding/soil/pkg/linkace/response"

func NewSlice(
	v []response.Link,
	host string,
) []*Link {
	var result []*Link

	for _, e := range v {
		result = append(result, New(e, host))
	}

	return result
}
