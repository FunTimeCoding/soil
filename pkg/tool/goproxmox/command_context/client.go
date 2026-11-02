package command_context

import "github.com/funtimecoding/soil/pkg/tool/goproxmoxd/face"

func (c *Context) Client() face.Client {
	return c.client
}
