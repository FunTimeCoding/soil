package chat_response

import (
	"github.com/funtimecoding/soil/pkg/console/status"
	"github.com/funtimecoding/soil/pkg/console/status/option"
)

func (r *Response) Format(f *option.Format) string {
	return status.New(f).String(
		r.Message.Role,
		r.Message.Content,
	).RawList(r.Raw).Format()
}
