package request_server

import "github.com/funtimecoding/soil/pkg/gitlab/merge_request"

func Identifiers(v []*merge_request.Request) []int64 {
	var result []int64

	for _, r := range v {
		result = append(result, r.Identifier)
	}

	return result
}
