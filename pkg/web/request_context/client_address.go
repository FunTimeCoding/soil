package request_context

import "github.com/funtimecoding/soil/pkg/web"

func (c *Context) ClientAddress() string {
	return web.ClientAddress(c.request)
}
