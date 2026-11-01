package command_context

import "github.com/funtimecoding/soil/pkg/tool/goclauded/face"

func (c *Context) LongClient() face.Client {
	return c.longClient
}
