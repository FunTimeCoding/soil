package search_result

import "github.com/funtimecoding/soil/pkg/atlassian/confluence/basic/response"

type Result struct {
	Name string
	Raw  *response.Result
}
