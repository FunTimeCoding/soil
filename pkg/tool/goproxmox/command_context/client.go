package command_context

import "github.com/funtimecoding/soil/pkg/tool/goproxmoxd/client"

func (c *Context) Client() *client.Client {
	return c.client
}
