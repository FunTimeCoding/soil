package page

import "github.com/funtimecoding/soil/pkg/atlassian/confluence/basic/response"

func NewSlice(
	v []*response.Page,
	host string,
) []*Page {
	var result []*Page

	for _, c := range v {
		result = append(result, New(c, host))
	}

	return result
}
