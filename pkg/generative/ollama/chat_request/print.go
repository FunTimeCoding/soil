package chat_request

import "github.com/funtimecoding/soil/pkg/console"

func (r *Request) Print() {
	for _, m := range r.request.Messages {
		console.Format("%s: %s\n", m.Role, m.Content)
	}
}
