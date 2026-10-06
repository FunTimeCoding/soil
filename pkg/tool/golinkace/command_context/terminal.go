package command_context

import "github.com/funtimecoding/soil/pkg/terminal"

func (c *Context) Terminal() *terminal.Terminal {
	return c.terminal
}
