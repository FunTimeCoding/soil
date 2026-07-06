package search_result

import "github.com/funtimecoding/soil/pkg/atlassian/confluence/basic/response"

func NewSlice(v []*response.Result) []*Result {
	var result []*Result

	for _, c := range v {
		result = append(result, New(c))
	}

	return result
}
