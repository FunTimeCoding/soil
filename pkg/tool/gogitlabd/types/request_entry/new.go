package request_entry

import "github.com/funtimecoding/soil/pkg/gitlab/merge_request"

func New(
	instance string,
	r *merge_request.Request,
) *Entry {
	return &Entry{Instance: instance, Request: r}
}
