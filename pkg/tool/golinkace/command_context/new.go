package command_context

import "github.com/funtimecoding/soil/pkg/terminal"

func New(t *terminal.Terminal) *Context {
	return &Context{terminal: t}
}
