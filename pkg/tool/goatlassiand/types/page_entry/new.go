package page_entry

import "github.com/funtimecoding/soil/pkg/atlassian/confluence/basic/response"

func New(page *response.Page) *Entry {
	return &Entry{Page: page}
}
