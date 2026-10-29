package page_entry

import "github.com/funtimecoding/soil/pkg/atlassian/confluence/basic/response"

type Entry struct {
	Page    *response.Page
	Draft   *response.Page
	Deleted bool
}
