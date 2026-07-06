package mock_client

import "github.com/funtimecoding/soil/pkg/atlassian/confluence/basic/response"

type entry struct {
	page    *response.Page
	draft   *response.Page
	deleted bool
}
