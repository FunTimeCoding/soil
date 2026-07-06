package space

import "github.com/funtimecoding/soil/pkg/atlassian/confluence/basic/response"

func NewSlice(v []*response.Space) []*Space {
	var result []*Space

	for _, c := range v {
		result = append(result, New(c))
	}

	return result
}
