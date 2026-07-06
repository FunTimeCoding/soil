package request_context

import "github.com/funtimecoding/soil/pkg/web"

func (c *Context) Write(
	code int,
	s string,
) {
	web.Write(c.writer, code, s)
}
