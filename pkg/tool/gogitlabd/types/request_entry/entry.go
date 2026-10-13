package request_entry

import "github.com/funtimecoding/soil/pkg/gitlab/merge_request"

type Entry struct {
	Instance string
	Request  *merge_request.Request
}
