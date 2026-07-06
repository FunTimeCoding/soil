package command_context

import "github.com/funtimecoding/soil/pkg/tool/goproxmoxd/client"

func (c *Context) Initialize(instance string) {
	c.client = client.NewEnvironment(instance)
}
